package container

import (
	"github.com/kearth/klib/kctx"
)

// // Module 模块接口
// type Module interface {
// 	Register() error
// 	Name() string
// }

// Component 组件接口
type Component interface {
	Unit
	Init(ctx kctx.Context) error
}

// Step 步骤
type Step struct {
	// BaseObject
	unit
	s func(ctx kctx.Context) error
}

// NewStep 创建步骤
func NewStep(name string, s func(ctx kctx.Context) error) Step {
	return Step{
		// BaseObject: BaseObject{Name: name},
		s: s,
	}
}

// Run 运行步骤
func (s Step) Run(ctx kctx.Context) error {
	return s.s(ctx)
}
