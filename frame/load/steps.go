package load

import (
	"context"

	"github.com/kearth/tea/frame/arg"
	"github.com/kearth/tea/frame/container"
)

type BaseFrameStep struct {
	container.BaseObject
}

func (f *BaseFrameStep) Type() StepType {
	return FrameStep
}

func (f *BaseFrameStep) Run(ctx context.Context, param ...arg.Arg) error {
	return nil
}

type BaseUserStep struct {
	container.BaseObject
}

func (u *BaseUserStep) Type() StepType {
	return UserStep
}

func (u *BaseUserStep) Run(ctx context.Context, param ...arg.Arg) error {
	return nil
}
