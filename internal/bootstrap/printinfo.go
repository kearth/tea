package bootstrap

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/frame/tlog"
	"github.com/kearth/tea/frame/utils"
)

var (
	// 实现接口
	_ container.Component = (*PrintInfo)(nil)

	// 单例
	printInfoInstance = &PrintInfo{}
)

// PrintInfo 打印信息
type PrintInfo struct {
	container.BaseObject
}

// LoadPrintInfo 获取单例
func LoadPrintInfo() *PrintInfo {
	return printInfoInstance
}

// Init 初始化
func (p *PrintInfo) Init(ctx tctx.Context) error {
	p.SetName("Print")
	envInfo := LoadEnvInfo()
	tlog.Info(ctx, "==================== Running Info Begin ====================")
	tlog.Info(ctx, utils.SPF("Version [%v]", base.Version()))
	tlog.Info(ctx, utils.SPF("Mode [%v]", envInfo.Mode))
	tlog.Info(ctx, utils.SPF("OS [%v]", envInfo.OS))
	tlog.Info(ctx, utils.SPF("OSVersion [%s]", envInfo.SystemVersion))
	tlog.Info(ctx, utils.SPF("Arch [%v]", envInfo.Arch))
	tlog.Info(ctx, utils.SPF("CPU [%d]", envInfo.CPU))
	tlog.Info(ctx, utils.SPF("Hostname [%s]", envInfo.Hostname))
	tlog.Info(ctx, utils.SPF("Pid [%d]", envInfo.PID))
	tlog.Info(ctx, utils.SPF("RootDir [%s]", envInfo.RootDir))
	tlog.Info(ctx, utils.SPF("ResourcesDir [%s]", envInfo.ResourcesDir))
	tlog.Info(ctx, utils.SPF("ServerVersion [%s]", envInfo.ServerVersion))
	tlog.Notice(ctx, utils.SPF("Server [%s] is starting...", gstr.UcFirst(envInfo.ServerName)))
	tlog.Notice(ctx, utils.SPF("Listening on [%s]", envInfo.Address))
	tlog.Info(ctx, "==================== Running Info End ====================")
	return nil
}
