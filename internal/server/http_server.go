package server

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/t"
	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/frame/tlog"
	"github.com/kearth/tea/frame/utils"
)

// Server 服务接口
var _ container.Server = (*HTTPServer)(nil)

// Register 注册服务
var _ = t.RegisterServer(base.DefaultServer, NewHTTPServer())

// RedocUI Redoc UI模板
const RedocUI = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Redoc UI</title>
  </head>
  <body>
    <redoc spec-url="{SwaggerUIDocUrl}"></redoc>
    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
  </body>
</html>`

// ServerPackage 服务包
type ServerPackage struct {
	container.BaseObject
}

// Register 注册服务包
func (s *ServerPackage) Register() error {
	s.SetName("ServerPackage")
	return nil
}

// HTTPConfig HTTP服务配置
type HTTPConfig struct {
	Router    container.Router     `json:"router"`    // 路由
	Resources []container.Resource `json:"resources"` // 资源
	Listeners []container.Listener `json:"listeners"` // 监听器
}

// HTTPServer HTTP服务
type HTTPServer struct {
	container.BaseObject
	Serv      *ghttp.Server
	Port      int                  // 端口
	Router    container.Router     // 路由
	Resources []container.Resource // 资源
	Listeners []container.Listener // 监听器
}

// New 创建实例
func NewHTTPServer() *HTTPServer {
	hs := &HTTPServer{
		Serv: g.Server(),
	}
	hs.SetName("HTTPServer")
	return hs
}

// Parse 解析配置
func (h *HTTPConfig) Parse(ctx tctx.Context) error {
	routerName := env.Cfg().MustGet(ctx, "server.router", "HTTPRouter").String()
	h.Router = t.GetRouter(routerName)
	listenerName := env.Cfg().MustGet(ctx, "server.listeners").Strings()
	for _, l := range listenerName {
		h.Listeners = append(h.Listeners, t.GetListener(l))
	}
	resourceName := env.Cfg().MustGet(ctx, "server.resources").Strings()
	for _, r := range resourceName {
		h.Resources = append(h.Resources, t.GetResource(r))
	}
	return nil
}

// Set 设置服务
func (h *HTTPServer) Set(ctx tctx.Context) error {
	h.Port = env.Port()
	cfg := &HTTPConfig{}
	err := cfg.Parse(ctx)
	if err != nil {
		return err
	}
	h.Router = cfg.Router
	h.Resources = cfg.Resources
	h.Listeners = cfg.Listeners
	return nil
}

// Start 启动服务
func (h *HTTPServer) Start(ctx tctx.Context) error {
	// 服务配置
	var config ghttp.ServerConfig

	// debug模式
	if env.IsDebug() {
		// 全局返回值
		oa := h.Serv.GetOpenApi()
		oa.Config.CommonResponse = base.Output{}
		oa.Config.CommonResponseDataField = "Data"
		config = ghttp.ServerConfig{
			Graceful:          true,                     // 优雅重启
			PProfEnabled:      true,                     // 开启pprof
			OpenApiPath:       "/api.json",              // 设置openapi路径
			SwaggerPath:       "/swagger",               // 设置swagger路径
			SwaggerUITemplate: RedocUI,                  // 设置swagger模板
			Address:           utils.SPF(":%d", h.Port), // 设置监听端口
			Logger:            tlog.Logger(),            // 设置日志
			DumpRouterMap:     true,                     // 打印路由信息
			LogStdout:         true,                     // 输出日志到控制台
			ErrorStack:        true,                     // 打印错误堆栈
		}
	} else {
		config = ghttp.ServerConfig{
			Graceful:   true,                     // 优雅重启
			Address:    utils.SPF(":%d", h.Port), // 设置监听端口
			Logger:     tlog.Logger(),            // 设置日志
			LogStdout:  true,                     // 输出日志到控制台
			ErrorStack: true,                     // 打印错误堆栈
		}
	}

	_ = h.Serv.SetConfig(config) // 设置配置
	h.Router.Register(h)         // 注册路由
	h.Serv.Run()                 // 启动服务
	return nil
}

// Stop 停止服务
func (h *HTTPServer) Stop(ctx tctx.Context) error {
	// 关闭监听器
	for _, v := range h.Listeners {
		_ = v.Close()
	}
	// 释放资源
	for _, v := range h.Resources {
		_ = v.Release(ctx)
	}
	// 关闭http服务
	return h.Serv.Shutdown()
}
