package httpserver

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kearth/klib/kunit"
	"github.com/kearth/tea/frame/server"
)

// Method HTTP 方法
type Method string

// Request HTTP 请求
type Request = ghttp.Request

const (

	// HTTP 方法
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
	TRACE   Method = "TRACE"
	CONNECT Method = "CONNECT"
	PATCH   Method = "PATCH"
	ALL     Method = "ALL"
)

func init() {
	// 注册 HTTP 路由
	server.RegisterRouter(server.HTTPRouterName, &HTTPRouter{
		Unit:         kunit.NewUnit(server.HTTPRouterName).SetRole(kunit.RoleRouter),
		groupsPrefix: "/",
		binds:        []any{},
		middlewares:  []Middleware{},
		groups:       []*Group{},
	})
}
