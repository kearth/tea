package tea

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/frame/container"
	tt "github.com/kearth/tea/frame/t"
	"github.com/kearth/tea/frame/tlog"
	"github.com/kearth/tea/internal/bootstrap"
	"github.com/kearth/tea/internal/cmd"
	"github.com/kearth/tea/internal/server"
)

// Tea 茶
type Tea struct {
	container.Unit
	RootCtx kctx.Context
}

// New
func New(ctx kctx.Context) *Tea {
	t := &Tea{
		Unit:    container.NewUnit("Tea"),
		RootCtx: ctx,
	}
	return t
}

// Run
func (t *Tea) Run() {
	ctx := t.RootCtx
	// 开始
	tlog.Info(ctx, "******************** Begin ********************")

	// 注册服务包
	if err := t.registerPackages(ctx); err != nil {
		tlog.Error(ctx, err.Error())
		return
	}

	// 安装组件 (按顺序安装)
	if err := t.installComponents(ctx); err != nil {
		tlog.Error(ctx, err.Error())
		return
	}

	tlog.Info(ctx, "******************** End ********************")
}

// registerPackages 注册服务包
func (t *Tea) registerPackages(ctx kctx.Context) error {
	packages := []container.Module{
		&server.ServerPackage{},
	}

	for _, p := range packages {
		if err := tt.AddModule(p); err != nil {
			tlog.Error(ctx, fmt.Sprintf("[Module][%s] add failed, err:%e", p.Name(), err))
			return err // 发生错误时停止
		}
		tlog.Notice(ctx, fmt.Sprintf("[Module][%s] add success", p.Name()))
	}
	return nil
}

// installComponents 安装组件
func (t *Tea) installComponents(ctx kctx.Context) error {
	components := []container.Component{
		bootstrap.LoadEnvInfo(),   // 环境
		cmd.Instance(),            // 命令行
		bootstrap.LoadPrintInfo(), // 打印
		bootstrap.Loads(),         // 加载器
	}

	// 安装组件并处理错误
	return bootstrap.Install(ctx, components, func(s string) {
		tlog.Notice(ctx, s)
	})
}
