package container

import "github.com/kearth/tea/frame/tctx"

// Package 包接口
type Package interface {
	Register() error
	Name() string
}

// Component 组件接口
type Component interface {
	Object
	Init(ctx tctx.Context) error
}

// Step 步骤
type Step struct {
	BaseObject
	s func(ctx tctx.Context) error
}

// NewStep 创建步骤
func NewStep(name string, s func(ctx tctx.Context) error) Step {
	return Step{
		BaseObject: BaseObject{_name: name},
		s:          s,
	}
}

// Run 运行步骤
func (s Step) Run(ctx tctx.Context) error {
	return s.s(ctx)
}
