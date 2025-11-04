package bootstrap

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
	"github.com/kearth/klib/klog"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/server"
)

var (
	loadInstance = &Load{}
)

// Step 步骤
type Step struct {
	container.Unit
	s func(ctx kctx.Context) error
}

// NewStep 创建步骤
func NewStep(name string, s func(ctx kctx.Context) error) Step {
	unit := container.NewUnit(name)
	return Step{
		Unit: unit,
		s:    s,
	}
}

// Run 运行步骤
func (s Step) Run(ctx kctx.Context) error {
	return s.s(ctx)
}

// Load 加载器
type Load struct {
	container.Unit
}

// LoadInstance
func LoadInstance() *Load {
	return loadInstance
}

// Setup
func (l *Load) Setup(ctx kctx.Context) kerr.Error {
	// Load
	unit := container.NewUnit("Load")
	unit.SetRole(container.RoleComponent)
	l.Unit = unit

	var err error
	// 解析步骤
	s := server.GetServers()
	for _, serv := range s {
		for _, step := range []Step{
			NewStep("Init", serv.Init),
			NewStep("Starting", serv.Start),
			NewStep("Stopping", serv.Stop),
		} {
			klog.Print(ctx, fmt.Sprintf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> [Step][%s] success >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", step.Name()))
			if err = step.Run(ctx); err != nil {
				klog.Print(ctx, fmt.Sprintf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> [Step][%s] error:%e >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", step.Name(), err))
				return base.StepError.Wrap(err)
			}
		}
	}
	return nil
}
