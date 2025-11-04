package server

import (
	"github.com/kearth/tea/frame/container"
)

// HTTPRouter 实现 HTTP 路由接口
var _ container.Router = (*HTTPRouter)(nil)

// 注册 HTTP 路由
var _ = RegisterRouter("HTTPRouter", &HTTPRouter{})

// HTTPRouter HTTP 路由
type HTTPRouter struct {
	container.Unit
}

// Register 注册路由
func (h *HTTPRouter) Register(s container.Server) {
	h.Unit = container.NewUnit("HTTPRouter")
	h.Unit.SetRole(container.RoleRouter)
}
