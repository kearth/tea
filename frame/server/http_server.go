package server

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/tctx"
)

var _ Server = (*HTTPServer)(nil)

// HTTPServer HTTP服务
type HTTPServer struct {
	container.BaseObject
	Serv      *ghttp.Server
	Port      int        // 端口
	Router    Router     // 路由
	Resources []Resource // 资源
	Listeners []Listener // 监听器
}

// New 创建实例
func NewHTTPServer(port int) *HTTPServer {
	hs := &HTTPServer{
		Serv: g.Server(),
		Port: port,
	}
	hs.SetName("HTTPServer")
	return hs
}

// SetRouter 设置路由
func (h *HTTPServer) SetRouter(router Router) {
	h.Router = router
}

// SetListener 设置监听器
func (h *HTTPServer) SetListener(ctx tctx.Context, listener []Listener) {
	h.Listeners = listener
}

// AddResource 添加资源
func (h *HTTPServer) AddResource(resource Resource) {
	h.Resources = append(h.Resources, resource)
}

// LoadResource 加载资源
func (h *HTTPServer) LoadResource(ctx tctx.Context) error {
	return nil
}

// Start 启动服务
func (h *HTTPServer) Start(ctx tctx.Context) error {
	return nil
}

// Stop 停止服务
func (h *HTTPServer) Stop(ctx tctx.Context) error {
	return nil
}
