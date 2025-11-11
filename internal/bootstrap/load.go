package bootstrap

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
	"github.com/kearth/klib/klog"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server"
	"github.com/kearth/tea/frame/server/httpserver"
)

var (
	loadInstance = &Load{
		Unit: container.NewUnit("Load").SetRole(container.RoleComponent),
	}
)

// Step 步骤
type Step struct {
	container.Unit
	s func(ctx kctx.Context) kerr.Error
}

// NewStep 创建步骤
func NewStep(name string, s func(ctx kctx.Context) kerr.Error) Step {
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
	return nil
}

func (l *Load) PrintSuccess(ctx kctx.Context, step Step) {
	klog.ColorPrint(ctx, klog.Yellow, fmt.Sprintf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> [Step][%s][%s] success >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", step.Name(), step.Cost()))
}

func (l *Load) PrintError(ctx kctx.Context, step Step, err error) {
	klog.ColorPrint(ctx, klog.Red, fmt.Sprintf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> [Step][%s][%s] error:%e >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", step.Name(), step.Cost(), err))
}

func (l *Load) Run(ctx kctx.Context) error {
	var err error
	var s server.Server
	// 解析步骤
	switch env.GetServerType() {
	case base.ServerTypeHTTP:
		fallthrough
	default:
		err = httpserver.NewHTTPServer().Setup(ctx)
		if err != nil {
			return base.ServerSetupError.Wrap(err)
		}
		s = server.GetServer(server.HTTPServerName)
	}
	for _, step := range []Step{
		NewStep("Init", s.Init),
		NewStep("Running", s.Start),
		NewStep("Stopping", s.Stop),
	} {
		step.SetFn(func(ctx kctx.Context, input ...any) (output any, err kerr.Error) {
			if sErr := step.Run(ctx); sErr != nil {
				return nil, base.StepError.Wrap(sErr)
			}
			return nil, nil
		})

		if _, err = step.Call(ctx); err != nil {
			l.PrintError(ctx, step, err)
			return base.StepError.Wrap(err)
		}
		l.PrintSuccess(ctx, step)
	}
	return nil
}
