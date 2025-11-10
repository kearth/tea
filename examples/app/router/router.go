package router

import (
	v1 "example/local/app/api/v1"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server"
	"github.com/kearth/tea/frame/server/httpserver"
)

func LoadRouter(ctx kctx.Context) error {
	var router = server.GetRouter(ctx, env.GetServerType())
	if hr, ok := router.(*httpserver.HTTPRouter); ok {
		Router(hr)
	}
	return nil
}

// Router 注册路由
func Router(hr *httpserver.HTTPRouter) {
	hr.SetGroupsPrefix("/api")
	hr.AddBind(new(v1.Hello))
	hr.AddBind(new(v1.Welcome))
	hr.AddMiddleware(httpserver.MiddlewareHandlerResponse())
}
