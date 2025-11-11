package router

import (
	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server"
	"github.com/kearth/tea/frame/server/httpserver"
)

// LoadRouter 加载路由
func LoadRouter(ctx kctx.Context) error {
	var router = server.GetRouter(ctx, env.GetServerType())
	switch env.GetServerType() {
	case base.ServerTypeHTTP:
		if hr, ok := router.(*httpserver.HTTPRouter); ok {
			Router(hr)
		}
	}
	return nil
}
