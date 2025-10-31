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
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/t"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	// 实例
	envInfoInstance = &EnvInfo{}
)

// 环境
type EnvInfo struct {
	container.Unit
	Version       string
	ServerName    string `json:"server_name"`
	ServerVersion string `json:"server_version"`
	Mode          string `json:"mode"`
	OS            string
	Port          int    `json:"port"`
	IP            string `json:"ip"`
	Host          string `json:"host"`
	Address       string
	Arch          string
	SystemVersion string
	CPU           int
	PID           int
	Hostname      string
	RootDir       string `json:"root_dir"`
	ResourcesDir  string `json:"resources_dir"`
	Cfg           *gcfg.Config
}

// EnvInfoInstance 获取实例
func EnvInfoInstance() *EnvInfo {
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

// 初始化环境
func (e *EnvInfo) Setup(ctx kctx.Context) kerr.Error {
	// 创建单元
	unit := container.NewUnit("EnvInfo")
	unit.SetRole(container.RoleComponent)
	e.Unit = unit
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
	e.ResourcesDir = cfg.MustGet(ctx, "server.resources_dir", "./resources").String()
	e.Mode = cfg.MustGet(ctx, "server.mode", base.EnvModeNormal).String()
	e.IP = cfg.MustGet(ctx, "server.ip", "").String()
	e.Port = cfg.MustGet(ctx, "server.port", 8080).Int()
	e.Host = cfg.MustGet(ctx, "server.host", "localhost").String()

	if e.IP != "" {
		e.Address = fmt.Sprintf("%s:%d", e.IP, e.Port)
	} else {
		e.Address = fmt.Sprintf("%s:%d", e.Host, e.Port)
	}

	e.SystemVersion = e.getSystemVersion(ctx)
	e.OS = runtime.GOOS
	e.Arch = runtime.GOARCH
	e.CPU = runtime.NumCPU()
	e.Hostname, _ = os.Hostname()
	e.PID = os.Getpid()

	// 设置默认服务器
	defaultServer := cfg.MustGet(ctx, "framework.default_server", base.DefaultServer).String()
	klog.Print(ctx, "======================================== Env Info =========================================")
	klog.Print(ctx, fmt.Sprintf("Version:         [ %v ]", e.Version))
	klog.Print(ctx, fmt.Sprintf("Mode:            [ %v ]", e.Mode))
	klog.Print(ctx, fmt.Sprintf("OS:              [ %v ]", e.OS))
	klog.Print(ctx, fmt.Sprintf("OSVersion:       [ %s ]", e.SystemVersion))
	klog.Print(ctx, fmt.Sprintf("Arch:            [ %v ]", e.Arch))
	klog.Print(ctx, fmt.Sprintf("CPU:             [ %d ]", e.CPU))
	klog.Print(ctx, fmt.Sprintf("Hostname:        [ %s ]", e.Hostname))
	klog.Print(ctx, fmt.Sprintf("Pid:             [ %d ]", e.PID))
	klog.Print(ctx, fmt.Sprintf("RootDir:         [ %s ]", e.RootDir))
	klog.Print(ctx, fmt.Sprintf("ResourcesDir:    [ %s ]", e.ResourcesDir))
	klog.Print(ctx, fmt.Sprintf("ServerVersion:   [ %s ]", e.ServerVersion))
	klog.Notice(ctx, fmt.Sprintf("Server:          [ %s ] is starting...", gstr.UcFirst(e.ServerName)))
	klog.Notice(ctx, fmt.Sprintf("Listening on:    [ %s ]", e.Address))
	return t.SetServer(defaultServer)
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
