package load

import (
	"context"

	"github.com/kearth/tea/frame/arg"
)

// 流程类型
type StepType string

const (
	FrameStep StepType = "frame_step" // 框架流程
	UserStep  StepType = "user_step"  // 用户流程
)

// 执行流程
type Step interface {
	Name() string                                    // 流程名称
	Type() StepType                                  // 流程类型
	Run(ctx context.Context, param ...arg.Arg) error // 执行流程
}
