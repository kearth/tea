package bootstrap

import (
	"fmt"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/tlog"
)

var (
	// 实现接口
	_ container.Component = (*PrintInfo)(nil)

	// 单例
	printInfoInstance = &PrintInfo{}
)

// PrintInfo 打印信息
type PrintInfo struct {
	container.Unit
}

// LoadPrintInfo 获取单例
func LoadPrintInfo() *PrintInfo {
	return printInfoInstance
}

// Init 初始化
func (p *PrintInfo) Init(ctx kctx.Context) error {
	p.Unit = container.NewUnit("PrintInfo")
	envInfo := LoadEnvInfo()
	tlog.Info(ctx, "==================== Running Info Begin ====================")
	// tlog.Info(ctx, fmt.Sprintf("Version [%v]", base.Version()))
	tlog.Info(ctx, fmt.Sprintf("Mode [%v]", envInfo.Mode))
	tlog.Info(ctx, fmt.Sprintf("OS [%v]", envInfo.OS))
	tlog.Info(ctx, fmt.Sprintf("OSVersion [%s]", envInfo.SystemVersion))
	tlog.Info(ctx, fmt.Sprintf("Arch [%v]", envInfo.Arch))
	tlog.Info(ctx, fmt.Sprintf("CPU [%d]", envInfo.CPU))
	tlog.Info(ctx, fmt.Sprintf("Hostname [%s]", envInfo.Hostname))
	tlog.Info(ctx, fmt.Sprintf("Pid [%d]", envInfo.PID))
	tlog.Info(ctx, fmt.Sprintf("RootDir [%s]", envInfo.RootDir))
	tlog.Info(ctx, fmt.Sprintf("ResourcesDir [%s]", envInfo.ResourcesDir))
	tlog.Info(ctx, fmt.Sprintf("ServerVersion [%s]", envInfo.ServerVersion))
	tlog.Notice(ctx, fmt.Sprintf("Server [%s] is starting...", gstr.UcFirst(envInfo.ServerName)))
	tlog.Notice(ctx, fmt.Sprintf("Listening on [%s]", envInfo.Address))
	tlog.Info(ctx, "==================== Running Info End ====================")
	return nil
}
