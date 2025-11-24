package load

import (
	apiv1 "example/local/app/api/demo/v1"
	pagev1 "example/local/app/page/demo/v1"
	pagev2 "example/local/app/page/demo/v2"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server"
	"github.com/kearth/tea/frame/server/httpserver"
)

// LoadRouter 加载路由
func LoadRouter(ctx kctx.Context) (string, error) {
	var router = server.GetRouter(ctx, env.GetServerType())
	switch env.GetServerType() {
	case base.ServerTypeHTTP:
		if hr, ok := router.(*httpserver.HTTPRouter); ok {
			HttpRouterRegister(hr)
		}
	}
	return "HttpRouter", nil
}

// HttpRouterRegister 注册 HTTP 路由
func HttpRouterRegister(hr *httpserver.HTTPRouter) {

	// 全局中间件
	hr.AddMiddleware(httpserver.MiddlewareHandlerResponse)

	// 静态文件路由
	hr.AddStaticPath("/tailwind.css", "css/demo/tailwind.css")

	// API 路由组
	hr.Group("/api/demo", func(group *httpserver.Group) {
		group.POST("v1/say", new(apiv1.Hello).Hello)
		group.POST("v1/download", new(apiv1.File).Image)
		group.POST("v1/stream", new(apiv1.Stream).Chat)
	})

	// 页面 路由组
	hr.Group("/page/demo", func(group *httpserver.Group) {
		group.GET("v1/welcome", new(pagev1.Welcome).Welcome)
		group.GET("v2/welcome", new(pagev2.Welcome).Welcome)
		group.GET("v1/stream", new(pagev1.Stream).Chat)
	})

}
