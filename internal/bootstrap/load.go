package bootstrap

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
	"github.com/kearth/klib/klog"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/t"
)

var (
	// instance 实例
	loadInstance = &Load{}
)

// Load 加载器
type Load struct {
	container.Unit
}

// Instance 获取实例
func Loads() *Load {
	return loadInstance
}

// Init 初始化
func (l *Load) Setup(ctx kctx.Context) kerr.Error {
	l.Unit = container.NewUnit("Load")
	var err error
	// 解析步骤
	s := t.GetServer()
	for _, step := range []container.Step{
		container.NewStep("Set", s.Set),
		container.NewStep("Start", s.Start),
		container.NewStep("Stop", s.Stop),
	} {
		if err = step.Run(ctx); err != nil {
			klog.Error(ctx, fmt.Sprintf("[Step][%s] error:%e", step.Name(), err))
			return base.StepError.Wrap(err)
		}
		klog.Notice(ctx, fmt.Sprintf("[Step][%s] success", step.Name()))
	}
	return nil
}
