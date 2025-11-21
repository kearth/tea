package httpserver

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kearth/klib/kunit"
	"github.com/kearth/tea/frame/server"
)

// HTTPRouter 实现 HTTP 路由接口
var _ server.Router = (*HTTPRouter)(nil)

// HTTPRouter HTTP 路由
type HTTPRouter struct {
	kunit.Unit
	groupsPrefix string
	binds        []any
	middlewares  []Middleware
	groups       []*Group
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

// Group 创建路由组
func (r *HTTPRouter) Group(prefix string, f func(group *Group)) {
	g := &Group{
		prefix:      prefix,
		binds:       map[string]any{},
		middlewares: []Middleware{},
	}
	f(g)
	r.groups = append(r.groups, g)
}
