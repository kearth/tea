package op

import (
	"context"

	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/tlog"
)

// 流程类型
type StepType string

const (
	FrameStep StepType = "frame_step" // 框架流程
	UserStep  StepType = "user_step"  // 用户流程
)

// 执行流程
type Step interface {
	container.Object
	Type() StepType                                        // 流程类型
	Run(ctx context.Context, param ...container.Arg) error // 执行流程
}

// BaseFrameStep 框架流程
type BaseFrameStep struct {
	container.BaseObject
}

func (f *BaseFrameStep) Type() StepType {
	return FrameStep
}

func (f *BaseFrameStep) Run(ctx context.Context, param ...container.Arg) error {
	return nil
}

// BaseUserStep 用户流程
type BaseUserStep struct {
	container.BaseObject
}

func (u *BaseUserStep) Type() StepType {
	return UserStep
}

func (u *BaseUserStep) Run(ctx context.Context, param ...container.Arg) error {
	return nil
}

// LoadAndRun 加载并执行流程
func LoadAndRun(ctx context.Context, steps []Step) error {
	if len(steps) == 0 {
		return nil
	}
	for _, step := range steps {
		if err := step.Run(ctx); err != nil {
			tlog.Error(ctx, "LoadErr", step.Name(), "err", err)
			return err
		}
	}
	return nil
}
