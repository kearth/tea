package server

import (
	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
	"github.com/kearth/klib/kutil"
	"github.com/kearth/tea/frame/base"
)

var (
	// 服务实例
	serverMap    = map[string]Server{}
	resourcesMap = map[string]Resource{}
	listenersMap = map[string]Listener{}
	routerMap    = map[string]Router{}
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

// Resource 资源接口
type Resource interface {
	Load(ctx kctx.Context) error    // 加载资源
	Release(ctx kctx.Context) error // 释放资源
}

// Listener 监听器接口
type Listener interface {
	Listen(ctx kctx.Context) error // 监听资源
	Close() error                  // 关闭资源
}

// GetServer 获取服务实例
func GetServers() map[string]Server {
	return serverMap
}

// RegisterServerConfig 注册服务配置
func RegisterServer(key string, s Server) error {
	serverMap[key] = s
	return nil
}

// RegisterResource 注册资源
func RegisterResource(key string, r Resource) error {
	resourcesMap[key] = r
	return nil
}

// RegisterListener 注册监听器
func RegisterListener(key string, l Listener) error {
	listenersMap[key] = l
	return nil
}

// RegisterRouter 注册路由
func RegisterRouter(key string, r Router) error {
	routerMap[key] = r
	return nil
}

// GetRouter 获取路由
func GetRouter(key string) Router {
	r := routerMap[key]
	kutil.If[func()](r == nil, func() {
		panic(base.RouterNotFound)
	})
	return r
}

// GetResource 获取资源
func GetResource(key string) Resource {
	r := resourcesMap[key]
	kutil.If[func()](r == nil, func() {
		panic(base.ResourceNotFound)
	})
	return r
}

// GetListener 获取监听器
func GetListener(key string) Listener {
	l := listenersMap[key]
	kutil.If[func()](l == nil, func() {
		panic(base.ListenerNotFound)
	})
	return l
}
