package httpserver

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
	"github.com/kearth/klib/klog"
	"github.com/kearth/klib/kunit"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server"
)

// Server 服务接口
var _ server.Server = (*HTTPServer)(nil)

// HTTPServer HTTP服务
type HTTPServer struct {
	kunit.Unit
	Serv   *ghttp.Server
	Port   int           // 端口
	Router server.Router // 路由
}

// New 创建实例
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		Unit: kunit.NewUnit(server.HTTPServerName).SetRole(kunit.RoleServer),
		Serv: g.Server(),
	}
}

// Register 注册服务包
func (s *HTTPServer) Setup(ctx kctx.Context) kerr.Error {
	// 注册服务
	server.RegisterServer(server.HTTPServerName, s)
	return nil
}

// Set 设置服务
func (h *HTTPServer) Init(ctx kctx.Context) kerr.Error {
	h.Port = env.Port()
	h.Router = server.GetRouter(ctx, env.GetServerType())
	h.Router.Register(h) // 注册路由
	return nil
}

// Start 启动服务
func (h *HTTPServer) Start(ctx kctx.Context) kerr.Error {
	// 服务配置
	var config ghttp.ServerConfig

	// debug模式
	if env.IsDebug() {
		// 全局返回值
		oa := h.Serv.GetOpenApi()
		oa.Config.CommonResponse = base.Output{}
		oa.Config.CommonResponseDataField = "Data"
		config = ghttp.ServerConfig{
			Graceful:          true,                       // 优雅重启
			PProfEnabled:      true,                       // 开启pprof
			OpenApiPath:       "/api.json",                // 设置openapi路径
			SwaggerPath:       "/swagger",                 // 设置swagger路径
			SwaggerUITemplate: server.RedocUI,             // 设置swagger模板
			Address:           fmt.Sprintf(":%d", h.Port), // 设置监听端口
			Logger:            klog.Logger(),              // 设置日志
			DumpRouterMap:     true,                       // 打印路由信息
			LogStdout:         true,                       // 输出日志到控制台
			ErrorStack:        true,                       // 打印错误堆栈
		}
	} else {
		config = ghttp.ServerConfig{
			Graceful:      true,                       // 优雅重启
			Address:       fmt.Sprintf(":%d", h.Port), // 设置监听端口
			Logger:        klog.Logger(),              // 设置日志
			LogStdout:     true,                       // 输出日志到控制台
			ErrorStack:    true,                       // 打印错误堆栈
			DumpRouterMap: true,                       // 打印路由信息
		}
	}

	h.Serv.SetConfig(config) // 设置配置
	h.Serv.Run()             // 启动服务
	return nil
}

// Stop 停止服务
func (h *HTTPServer) Stop(ctx kctx.Context) kerr.Error {
	// 关闭http服务
	if h.Serv.Status() == ghttp.ServerStatusRunning {
		err := h.Serv.Shutdown()
		if err != nil {
			return base.ServerShutdownError.Wrap(err)
		}
	}
	return nil
}
