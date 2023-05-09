package tea

import "context"

// IApp
type IApp interface {
	Init(ctx context.Context) error
	Start() error
	Shutdown(ctx context.Context, cancel context.CancelFunc) error
}

// Bootstrap
type Bootstrap func(ctx context.Context) error

// IServer
type IServer interface {
	IApp
	SetConf(path string) error
	SetBootstrap(bs Bootstrap) error
	SetRouter(ir IRouter) error
}

// IRouter
type IRouter interface {
	Group(pattern string) IRouter
	Use()
}

// BaseServer
type BaseServer struct{}
