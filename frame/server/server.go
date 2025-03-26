package server

import "github.com/kearth/tea/frame/tctx"

type Router interface {
}
type Resource interface {
}
type Listener interface {
}

// Server 服务器接口
type Server interface {
	SetRouter(router Router)
	SetListener(ctx tctx.Context, listener []Listener)
	AddResource(resource Resource)
	LoadResource(ctx tctx.Context) error
	Start(ctx tctx.Context) error
	Stop(ctx tctx.Context) error
}
