package server

import (
	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/t"
)

// HTTPResources 实现Resource接口
var _ container.Resource = (*HTTPResources)(nil)

// 注册HTTP资源
var _ = t.RegisterResource("HTTPResource", &HTTPResources{})

// HTTPResources HTTP资源
type HTTPResources struct {
	container.Unit
}

// Load 加载HTTP资源
func (h *HTTPResources) Load(ctx kctx.Context) error {
	h.SetName("HTTPResource")
	// TODO: 加载HTTP资源
	return nil
}

func (h *HTTPResources) Release(ctx kctx.Context) error {
	// TODO: 释放HTTP资源
	return nil
}
