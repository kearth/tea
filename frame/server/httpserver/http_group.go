package httpserver

import "fmt"

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
func (r *Group) Bind(method Method, path string, object any) {
	key := fmt.Sprintf("%s:%s", method, path)
	r.binds[key] = object
}
