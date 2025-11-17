package bootstrap

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
	"github.com/kearth/klib/klog"
	"github.com/kearth/klib/kunit"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/env"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	EnvName = "Env"
)

var (
	// 实例
	envInfoInstance = &EnvInfo{
		Unit: kunit.NewUnit(EnvName).SetRole(kunit.RoleComponent),
	}
)

// 环境
type EnvInfo struct {
	kunit.Unit
	env.Env
}

// Env 获取实例
func Env() *EnvInfo {
	return envInfoInstance
}

// ParseTOML 解析配置文件
func ParseTOML(path string) (*gcfg.Config, error) {
	if !gfile.Exists(path) {
		return nil, base.ConfigFileNotExists
	}
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(path)
	return g.Cfg(), nil
}

func (e *EnvInfo) SetVersion(version string) *EnvInfo {
	e.Version = version
	return e
}

// 初始化环境
func (e *EnvInfo) Setup(ctx kctx.Context) kerr.Error {
	// 解析配置文件
	cfg, err := ParseTOML(base.GetConfigPath())
	if err != nil {
		return base.ConfigFileNotExists.Wrap(err)
	}
	// 设置配置信息
	e.Cfg = cfg
	e.ServerName = cfg.MustGet(ctx, "server.name", "tea").String()
	e.ServerVersion = cfg.MustGet(ctx, "server.version", "1.0.0").String()
	e.RootDir = cfg.MustGet(ctx, "server.root_dir", gfile.Pwd()).String()
	if e.RootDir == "." || e.RootDir == "" || e.RootDir == "/" {
		e.RootDir = gfile.Pwd()
	}
	e.ResourcesDir = cfg.MustGet(ctx, "server.resources_dir", "./resources").String()
	e.Mode = cfg.MustGet(ctx, "server.mode", base.EnvModeNormal).String()
	// 调试模式
	if e.Mode != base.EnvModeDebug {
		e.Mode = base.EnvModeNormal
	}
	e.IP = cfg.MustGet(ctx, "server.ip", base.DefaultIP).String()
	e.Port = cfg.MustGet(ctx, "server.port", base.DefaultPort).Int()
	e.Address = fmt.Sprintf("%s:%d", e.IP, e.Port)
	e.ServerType = cfg.MustGet(ctx, "server.server_type", base.ServerTypeHTTP).String()
	e.SystemVersion = e.getSystemVersion(ctx)
	e.OS = runtime.GOOS
	e.Arch = runtime.GOARCH
	e.CPU = runtime.NumCPU()
	e.Hostname, _ = os.Hostname()
	e.PID = os.Getpid()

	klog.ColorPrint(ctx, klog.Cyan, "======================================== Env Info =========================================")
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("FrameworkVersion:         [ %v ]", e.Version))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("Mode:                     [ %v ]", e.Mode))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("OS:                       [ %v ]", e.OS))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("OSVersion:                [ %s ]", e.SystemVersion))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("Arch:                     [ %v ]", e.Arch))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("CPU:                      [ %d ]", e.CPU))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("Hostname:                 [ %s ]", e.Hostname))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("Pid:                      [ %d ]", e.PID))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("RootDir:                  [ %s ]", e.RootDir))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("ResourcesDir:             [ %s ]", e.ResourcesDir))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("ServerVersion:            [ %s ]", e.ServerVersion))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("ServerName:               [ %s ]", gstr.UcFirst(e.ServerName)))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("ServerType:               [ %s ]", e.ServerType))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("Listening:                [ %s ]", e.Address))
	env.Init(&e.Env)
	return nil
}

// 获取系统版本
func (e *EnvInfo) getSystemVersion(ctx context.Context) string {
	switch runtime.GOOS {
	case base.OSWindows:
		bs, err := gproc.ShellExec(ctx, "ver")
		if err == nil {
			reader := transform.NewReader(
				io.NopCloser(bytes.NewReader([]byte(bs))),
				simplifiedchinese.GBK.NewDecoder(),
			)
			newBytes, _ := io.ReadAll(reader)
			return gstr.TrimAll(string(newBytes))
		}
	case base.OSMac:
		bs, err := gproc.ShellExec(ctx, "sw_vers -productVersion")
		if err == nil {
			return gstr.TrimAll(bs)
		}
	}
	return ""
}
