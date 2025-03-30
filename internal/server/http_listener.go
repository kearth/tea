package server

import (
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/t"
	"github.com/kearth/tea/frame/tctx"
)

// HTTPListener 实现 HTTP 监听器接口
var _ container.Listener = (*HTTPListener)(nil)

// 注册 HTTP 监听器
var _ = t.RegisterListener("HTTPListener", &HTTPListener{})

// HTTPListener HTTP 监听器
type HTTPListener struct {
	container.BaseObject
}

// Listen 监听
func (l *HTTPListener) Listen(ctx tctx.Context) error {
	l.SetName("HTTPListener")
	return nil
}

// Close 关闭监听器
func (l *HTTPListener) Close() error {
	return nil
}
