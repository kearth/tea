package t

import (
	"github.com/kearth/klib/kerr"
	"github.com/kearth/klib/kutil"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/container"
)

var (
	// 服务实例
	serverInstance container.Server
	serverMap      = map[string]container.Server{}
	resourcesMap   = map[string]container.Resource{}
	listenersMap   = map[string]container.Listener{}
	routerMap      = map[string]container.Router{}
)

// SetServer 设置服务实例
func SetServer(name string) kerr.Error {
	if s, ok := serverMap[name]; ok {
		serverInstance = s
		return nil
	}
	return base.DefaultServerNotFound
}

// GetServer 获取服务实例
func GetServer() container.Server {
	return serverInstance
}

// RegisterServerConfig 注册服务配置
func RegisterServer(key string, s container.Server) error {
	serverMap[key] = s
	return nil
}

// RegisterResource 注册资源
func RegisterResource(key string, r container.Resource) error {
	resourcesMap[key] = r
	return nil
}

// RegisterListener 注册监听器
func RegisterListener(key string, l container.Listener) error {
	listenersMap[key] = l
	return nil
}

// RegisterRouter 注册路由
func RegisterRouter(key string, r container.Router) error {
	routerMap[key] = r
	return nil
}

// GetRouter 获取路由
func GetRouter(key string) container.Router {
	r := routerMap[key]
	kutil.If[func()](r == nil, func() {
		panic(base.RouterNotFound)
	})
	return r
}

// GetResource 获取资源
func GetResource(key string) container.Resource {
	r := resourcesMap[key]
	kutil.If[func()](r == nil, func() {
		panic(base.ResourceNotFound)
	})
	return r
}

// GetListener 获取监听器
func GetListener(key string) container.Listener {
	l := listenersMap[key]
	kutil.If[func()](l == nil, func() {
		panic(base.ListenerNotFound)
	})
	return l
}

// // AddModule 添加模块
// func AddModule(p container.Module) error {
// 	return p.Register()
// }
