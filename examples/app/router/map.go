package router

import (
	apiv1 "example/local/app/api/demo/v1"
	pagev1 "example/local/app/page/demo/v1"
	pagev2 "example/local/app/page/demo/v2"

	"github.com/kearth/tea/frame/server/httpserver"
)

// Router 注册路由
func Router(hr *httpserver.HTTPRouter) {
	// 全局中间件

	// API 路由组
	apiGroup := hr.NewGroup("/api")
	apiGroup.Use(httpserver.MiddlewareHandlerResponse)
	apiGroup.Bind(httpserver.GET, "v1/say", new(apiv1.Hello))
	apiGroup.Bind(httpserver.GET, "v1/download", new(apiv1.File))

	// 页面 路由组
	pageGroup := hr.NewGroup("/page")
	pageGroup.Use(httpserver.MiddlewareLog)
	pageGroup.Bind(httpserver.GET, "v1", new(pagev1.Welcome))
	pageGroup.Bind(httpserver.GET, "v2", new(pagev2.Welcome))

}
