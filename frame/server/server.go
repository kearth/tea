package server

import (
	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
	"github.com/kearth/klib/klog"
	"github.com/kearth/klib/kutil"
	"github.com/kearth/tea/frame/base"
)

const (
	HTTPRouterName = "HTTPRouter"
	HTTPServerName = "HTTPServer"
)

var (
	// 服务实例
	serverMap = map[string]Server{}
	routerMap = map[string]Router{}
)

// Server 服务器接口
type Server interface {
	Init(ctx kctx.Context) kerr.Error  // 设置服务器
	Start(ctx kctx.Context) kerr.Error // 启动服务器
	Stop(ctx kctx.Context) kerr.Error  // 停止服务器
}

// Router 路由接口
type Router interface {
	Register(s Server)
}

// GetServer 获取服务实例
func GetServer(key string) Server {
	return serverMap[key]
}

// RegisterServerConfig 注册服务配置
func RegisterServer(key string, s Server) {
	serverMap[key] = s
}

// RegisterRouter 注册路由
func RegisterRouter(key string, r Router) {
	routerMap[key] = r
}

// GetRouter 获取路由
func GetRouter(ctx kctx.Context, key string) Router {
	var r Router
	switch key {
	case base.ServerTypeHTTP:
		r = routerMap[HTTPRouterName]
	default:
		r = routerMap[HTTPRouterName]
	}
	kutil.If[func()](r == nil, func() {
		klog.Panic(ctx, base.RouterNotFound)
	})
	return r
}
