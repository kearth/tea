package tea

import (
	"github.com/kearth/tea/frame/container"
	tt "github.com/kearth/tea/frame/t"
	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/frame/tlog"
	"github.com/kearth/tea/frame/utils"
	"github.com/kearth/tea/internal/bootstrap"
	"github.com/kearth/tea/internal/cmd"
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
	err := bootstrap.Install(ctx, []container.Component{
		bootstrap.LoadEnvInfo(),   // 环境
		cmd.Instance(),            // 命令行
		bootstrap.LoadPrintInfo(), // 打印
		bootstrap.Loads(),         // 加载器
	}, func(s string) {
		tlog.Notice(ctx, s)
	})
	if err != nil {
		tlog.Error(ctx, err.Error())
	}

	tlog.Info(ctx, "******************** End ********************")
}
