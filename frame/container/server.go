package container

import "github.com/kearth/klib/kctx"

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

// ServerConfig 服务器配置接口
type ServerConfig interface {
	Router() Router        // 路由
	Resources() []Resource // 资源
	Listeners() []Listener // 监听器
}

// Server 服务器接口
type Server interface {
	Set(ctx kctx.Context) error   // 设置服务器
	Start(ctx kctx.Context) error // 启动服务器
	Stop(ctx kctx.Context) error  // 停止服务器
}
