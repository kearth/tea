package router

import (
	"example/local/api"

	"github.com/kearth/tea/frame/server"
	"github.com/kearth/tea/frame/server/httpserver"
)

func GetRouter() {

}

// Router 注册路由
func Router() {
	var router = server.GetRouter(httpserver.HTTPRouterName)
	if hr, ok := router.(*httpserver.HTTPRouter); ok {
		hr.SetGroupsPrefix("/api")
		hr.AddBind(new(api.Hello))

	}
}
