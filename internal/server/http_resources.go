package server

import (
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/t"
	"github.com/kearth/tea/frame/tctx"
)

// HTTPResources 实现Resource接口
var _ container.Resource = (*HTTPResources)(nil)

// 注册HTTP资源
var _ = t.RegisterResource("HTTPResource", &HTTPResources{})

// HTTPResources HTTP资源
type HTTPResources struct {
	container.BaseObject
}

// Load 加载HTTP资源
func (h *HTTPResources) Load(ctx tctx.Context) error {
	h.SetName("HTTPResource")
	// TODO: 加载HTTP资源
	return nil
}

func (h *HTTPResources) Release(ctx tctx.Context) error {
	// TODO: 释放HTTP资源
	return nil
}
