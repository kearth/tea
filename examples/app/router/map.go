package router

import (
	apiv1 "example/local/app/api/v1"
	pagev1 "example/local/app/page/v1"

	"github.com/kearth/tea/frame/server/httpserver"
)

// Router 注册路由
func Router(hr *httpserver.HTTPRouter) {
	// 全局中间件

	// API 路由组
	apiGroup := hr.NewGroup("/api")
	apiGroup.Use(httpserver.MiddlewareHandlerResponse)
	apiGroup.Bind(httpserver.GET, "v1", new(apiv1.Hello))

	// 页面 路由组
	pageGroup := hr.NewGroup("/page")
	pageGroup.Use(httpserver.MiddlewareLog)
	pageGroup.Bind(httpserver.GET, "v1", new(pagev1.Welcome))

}
