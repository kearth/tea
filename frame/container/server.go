package container

import "github.com/kearth/tea/frame/tctx"

// Router 路由接口
type Router interface {
	Object
	Register(s Server)
}

// Resource 资源接口
type Resource interface {
	Object
	Load(ctx tctx.Context) error    // 加载资源
	Release(ctx tctx.Context) error // 释放资源
}

// Listener 监听器接口
type Listener interface {
	Object
	Listen(ctx tctx.Context) error // 监听资源
	Close() error                  // 关闭资源
}

// ServerConfig 服务器配置接口
type ServerConfig interface {
	Object
	Router() Router        // 路由
	Resources() []Resource // 资源
	Listeners() []Listener // 监听器
}

// Server 服务器接口
type Server interface {
	Set(ctx tctx.Context) error   // 设置服务器
	Start(ctx tctx.Context) error // 启动服务器
	Stop(ctx tctx.Context) error  // 停止服务器
}
