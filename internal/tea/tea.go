package tea

import (
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/errs"
	tt "github.com/kearth/tea/frame/t"
	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/frame/tlog"
	"github.com/kearth/tea/frame/utils"
	"github.com/kearth/tea/internal/cmd"
	"github.com/kearth/tea/internal/envinfo"
	"github.com/kearth/tea/internal/load"
	"github.com/kearth/tea/internal/printinfo"
	"github.com/kearth/tea/internal/server"
)

// Tea 茶
type Tea struct {
	container.BaseObject
}

// New
func New() *Tea {
	t := new(Tea)
	t.BaseObject.SetName("Tea")
	return t
}

// Run
func (t *Tea) Run(ctx tctx.Context) {
	// 开始
	tlog.Info(ctx, "******************** Begin ********************")

	// 注册服务包
	for _, p := range []container.Package{
		&server.ServerPackage{},
	} {
		if err := tt.AddPackage(p); err != nil {
			tlog.Error(ctx, utils.SPF("[Package][%s] add failed, err:%e", p.Name(), err))
		} else {
			tlog.Notice(ctx, utils.SPF("[Package][%s] add success", p.Name()))
		}
	}

	// 安裝组件 (按顺序安装)
	err := Install(ctx, []container.Component{
		envinfo.Instance(),   // 环境
		cmd.Instance(),       // 命令行
		printinfo.Instance(), // 打印
		load.Instance(),      // 加载器
	}, func(s string) {
		tlog.Notice(ctx, s)
	})
	if err != nil {
		tlog.Error(ctx, err.Error())
	}

	tlog.Info(ctx, "******************** End ********************")
}

// Install 安装组件
func Install(ctx tctx.Context, c []container.Component, sf func(string)) error {
	for _, v := range c {
		if err := v.Init(ctx); err != nil {
			return errs.Wrap(err, utils.SPF("[%s] init error", v.Name()))
		}
		sf(utils.SPF("[Component][%s] init success", v.Name()))
	}
	return nil
}
