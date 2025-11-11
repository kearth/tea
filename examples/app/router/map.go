package router

import (
	v1 "example/local/app/api/v1"

	"github.com/kearth/tea/frame/server/httpserver"
)

// Router 注册路由
func Router(hr *httpserver.HTTPRouter) {
	hr.SetGroupsPrefix("/api")
	hr.AddBind(new(v1.Hello))
	hr.AddBind(new(v1.Welcome))
	hr.AddMiddleware(httpserver.MiddlewareHandlerResponse())
}
