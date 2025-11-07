package httpserver

import (
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/server"
)

// HTTPRouter 实现 HTTP 路由接口
var _ server.Router = (*HTTPRouter)(nil)

// Method HTTP 方法
type Method string

const (
	// HTTPRouterName HTTP 路由名称
	HTTPRouterName = "HTTPRouter"

	// HTTP 方法
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	ALL    Method = "ALL"
)

// 注册 HTTP 路由
var _ = server.RegisterRouter(HTTPRouterName, &HTTPRouter{
	Unit:         container.NewUnit(HTTPRouterName).SetRole(container.RoleRouter),
	groupsPrefix: "/",
	binds:        []any{},
	middlewares:  []Middleware{},
	groups:       []*Group{},
})

// Request HTTP 请求
type Request = ghttp.Request

// Middleware HTTP 中间件
type Middleware = func(r *Request)

// HTTPRouter HTTP 路由
type HTTPRouter struct {
	container.Unit
	groupsPrefix string
	binds        []any
	middlewares  []Middleware
	groups       []*Group
}

// Group 路由组
type Group struct {
	prefix      string
	binds       map[string]any
	middlewares []Middleware
}

// Register 注册路由
func (h *HTTPRouter) Register(s server.Server) {
	if hs, ok := s.(*HTTPServer); ok {
		hs.Serv.Group(h.groupsPrefix, func(group *ghttp.RouterGroup) {
			if len(h.middlewares) > 0 {
				group.Middleware(h.middlewares...)
			}
			if len(h.binds) > 0 {
				group.Bind(h.binds...)
			}
			if len(h.groups) > 0 {
				for _, gr := range h.groups {
					group.Group(gr.prefix, func(group *ghttp.RouterGroup) {
						if len(gr.middlewares) > 0 {
							group.Middleware(gr.middlewares...)
						}
						if len(gr.binds) > 0 {
							group.Map(gr.binds)
						}
					})
				}
			}
		})
	}
}

// SetGroupsPrefix 设置路由组前缀
func (r *HTTPRouter) SetGroupsPrefix(prefix string) {
	r.groupsPrefix = prefix
}

// AddMiddleware 添加中间件
func (r *HTTPRouter) AddMiddleware(middleware ...Middleware) {
	r.middlewares = append(r.middlewares, middleware...)
}

// AddBind 添加路由绑定对象
func (r *HTTPRouter) AddBind(bindObject ...any) {
	r.binds = append(r.binds, bindObject...)
}

// NewGroup 创建路由组
func (r *HTTPRouter) NewGroup(prefix string) *Group {
	g := &Group{
		prefix:      prefix,
		binds:       map[string]any{},
		middlewares: []Middleware{},
	}
	r.groups = append(r.groups, g)
	return g
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
