package httpserver

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

type Meta g.Meta

// Group 路由组
type Group struct {
	prefix      string
	binds       map[string]any
	middlewares []Middleware
}

// Use 添加中间件
func (r *Group) Use(middleware ...Middleware) {
	r.middlewares = append(r.middlewares, middleware...)
}

// Bind 注册路由
func (r *Group) Bind(method Method, prefix string, object any) {
	key := fmt.Sprintf("%s:%s", method, prefix)
	r.binds[key] = object
}

// GET 注册 GET 路由
func (r *Group) GET(prefix string, object any) {
	r.Bind(GET, prefix, object)
}

// POST 注册 POST 路由
func (r *Group) POST(prefix string, object any) {
	r.Bind(POST, prefix, object)
}

// PUT 注册 PUT 路由
func (r *Group) PUT(prefix string, object any) {
	r.Bind(PUT, prefix, object)
}

// DELETE 注册 DELETE 路由
func (r *Group) DELETE(prefix string, object any) {
	r.Bind(DELETE, prefix, object)
}

// HEAD 注册 HEAD 路由
func (r *Group) HEAD(prefix string, object any) {
	r.Bind(HEAD, prefix, object)
}

// OPTIONS 注册 OPTIONS 路由
func (r *Group) OPTIONS(prefix string, object any) {
	r.Bind(OPTIONS, prefix, object)
}

// TRACE 注册 TRACE 路由
func (r *Group) TRACE(prefix string, object any) {
	r.Bind(TRACE, prefix, object)
}

// CONNECT 注册 CONNECT 路由
func (r *Group) CONNECT(prefix string, object any) {
	r.Bind(CONNECT, prefix, object)
}

// PATCH 注册 PATCH 路由
func (r *Group) PATCH(prefix string, object any) {
	r.Bind(PATCH, prefix, object)
}

// ALL 注册 ALL 路由
func (r *Group) ALL(prefix string, object any) {
	r.Bind(ALL, prefix, object)
}
