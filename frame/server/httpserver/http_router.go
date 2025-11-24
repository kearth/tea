package httpserver

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/kearth/klib/kunit"
	"github.com/kearth/tea/frame/env"
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
	staticPaths  map[string]string
	staticRoot   string
}

// Register 注册路由
func (h *HTTPRouter) Register(s server.Server) {
	h.staticRoot = env.GetResourcesDir()
	if hs, ok := s.(*HTTPServer); ok {
		hs.Serv.Group(h.groupsPrefix, func(group *ghttp.RouterGroup) {
			if len(h.middlewares) > 0 {
				group.Middleware(h.middlewares...)
			}
			if len(h.binds) > 0 {
				group.Bind(h.binds...)
			}
			if len(h.staticPaths) > 0 {
				group.Map(bindStaticFile(h.staticRoot, h.staticPaths))
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

// SetStaticRoot 设置静态文件根目录
func (r *HTTPRouter) SetStaticRoot(root string) {
	if gfile.IsDir(root) {
		r.staticRoot = root
	}
}

// AddStaticPath 添加静态文件路由
func (r *HTTPRouter) AddStaticPath(uri, file string) {
	if r.staticPaths == nil {
		r.staticPaths = map[string]string{}
	}
	r.staticPaths[uri] = file
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
