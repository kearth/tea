package frame

import (
	"runtime"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kearth/tea/frame/cmd"
	"github.com/kearth/tea/frame/config"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/defined"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/op"
	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/frame/tlog"
	"github.com/kearth/tea/frame/utils"
)

// Tea 茶
type Tea struct {
	container.BaseObject
	version string
}

// Drink 喝
func (t *Tea) Drink(ctx tctx.Context) {
	var err error
	// 设置最大CPU核心数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化日志
	glog.SetDefaultHandler(tlog.DefaultHandler)

	// 开始
	tlog.Info(ctx, "******************** Begin ********************")

	// 指令系统
	new(cmd.Cmd).Init(ctx)

	// 初始化环境
	err = env.Instance().Init(ctx)
	if err != nil {
		tlog.Error(ctx, err.Error())
	}

	tlog.Info(ctx, utils.SPF("Version [%v]", t.version))
	tlog.Info(ctx, utils.SPF("Mode [%v]", env.Instance().Mode))
	tlog.Info(ctx, utils.SPF("OS [%v]", env.Instance().OS))
	tlog.Info(ctx, utils.SPF("OSVersion [%s]", env.Instance().SystemVersion))
	tlog.Info(ctx, utils.SPF("Arch [%v]", env.Instance().Arch))
	tlog.Info(ctx, utils.SPF("CPU [%d]", env.Instance().CPU))
	tlog.Info(ctx, utils.SPF("Hostname [%s]", env.Instance().Hostname))
	tlog.Info(ctx, utils.SPF("Pid [%d]", env.Instance().PID))
	tlog.Info(ctx, utils.SPF("RootDir [%s]", env.Instance().RootDir))
	tlog.Info(ctx, utils.SPF("ResourcesDir [%s]", env.Instance().ResourcesDir))
	tlog.Info(ctx, utils.SPF("ServerVersion [%s]", env.Instance().ServerVersion))
	tlog.Notice(ctx, utils.SPF("Env init done [%v]", utils.Condition(err == nil, "Success", "Fail")))
	tlog.Notice(ctx, utils.SPF("Server [%s] is starting...", gstr.UcFirst(env.Instance().ServerName)))
	tlog.Notice(ctx, utils.SPF("Listening on [%s]", env.Instance().Address))

	// 解析步骤
	steps := make([]op.Step, 0)
	config.ParseTOML(defined.LoadMapPath)
	// 执行步骤
	err = op.LoadAndRun(ctx, steps)
	if err != nil {
		tlog.Error(ctx, "Load some step error")
	}
	tlog.Info(ctx, "******************** End ********************")
}

// GetSomeTea 获取一些茶
func GetSomeTea(version string) *Tea {
	t := &Tea{version: version}
	// 设置名称
	t.BaseObject.SetName("Tea")
	return t
}
