package env

import (
	"bytes"
	"context"
	"io"
	"os"
	"runtime"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kearth/tea/frame/config"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/defined"
	"github.com/kearth/tea/frame/utils"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 实现组件接口
var _ container.Component = (*Env)(nil)

type ModeType string

func (m ModeType) String() string {
	return string(m)
}

type OSType string

func (m OSType) String() string {
	return string(m)
}

const (
	EnvModeNormal ModeType = "normal" // 正常模式
	EnvModeDebug  ModeType = "debug"  // 调试模式

	OSMac       OSType = "darwin"    // macOS
	OSLinux     OSType = "linux"     // Linux
	OSWindows   OSType = "windows"   // Windows
	OSAndroid   OSType = "android"   // Android
	OSFreeBSD   OSType = "freebsd"   // FreeBSD
	OSDragonfly OSType = "dragonfly" // Dragonfly
	OSNetBSD    OSType = "netbsd"    // NetBSD
	OSPlain     OSType = "plan9"     // Plan 9
	OSSolaris   OSType = "solaris"   // Solaris
	OSOpenBSD   OSType = "openbsd"   // OpenBSD
	OSUnknown   OSType = "unknown"   // 未知系统

)

// 实例
var instance *Env = &Env{
	Mode: EnvModeNormal,
}

// 环境
type Env struct {
	container.BaseObject
	Version       string
	ServerName    string   `json:"server_name"`
	ServerVersion string   `json:"server_version"`
	Mode          ModeType `json:"mode"`
	OS            OSType
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
}

func Instance() *Env {
	return instance
}

// 初始化环境
func (e *Env) Init(ctx context.Context) error {
	// 解析配置文件
	cfg, err := config.ParseTOML(defined.ConfigPath)
	if err != nil {
		return err
	}

	e.ServerName = cfg.MustGet(ctx, "server.name", "tea").String()
	e.ServerVersion = cfg.MustGet(ctx, "server.version", "1.0.0").String()
	e.RootDir = cfg.MustGet(ctx, "server.root_dir", gfile.Pwd()).String()
	e.ResourcesDir = cfg.MustGet(ctx, "server.resources_dir", "./resources").String()
	e.Mode = ModeType(cfg.MustGet(ctx, "server.mode", EnvModeNormal).String())
	e.IP = cfg.MustGet(ctx, "server.ip", "").String()
	e.Port = cfg.MustGet(ctx, "server.port", 8080).Int()
	e.Host = cfg.MustGet(ctx, "server.host", "localhost").String()

	if e.IP != "" {
		e.Address = utils.SPF("%s:%d", e.IP, e.Port)
	} else {
		e.Address = utils.SPF("%s:%d", e.Host, e.Port)
	}

	instance.SystemVersion = e.getSystemVersion(ctx)
	instance.OS = OSType(runtime.GOOS)
	instance.Arch = runtime.GOARCH
	instance.CPU = runtime.NumCPU()
	instance.Hostname, _ = os.Hostname()
	instance.PID = os.Getpid()
	instance.SetName("Env")
	return nil
}

// 获取系统版本
func (e *Env) getSystemVersion(ctx context.Context) string {
	switch runtime.GOOS {
	case OSWindows.String():
		bs, err := gproc.ShellExec(ctx, "ver")
		if err == nil {
			reader := transform.NewReader(
				io.NopCloser(bytes.NewReader([]byte(bs))),
				simplifiedchinese.GBK.NewDecoder(),
			)
			newBytes, _ := io.ReadAll(reader)
			return gstr.TrimAll(string(newBytes))
		}
	case OSMac.String():
		bs, err := gproc.ShellExec(ctx, "sw_vers -productVersion")
		if err == nil {
			return gstr.TrimAll(bs)
		}
	}
	return ""
}

func (e *Env) IsDebug() bool {
	return e.Mode == EnvModeDebug
}

func (e *Env) IsNormal() bool {
	return e.Mode == EnvModeNormal
}

func (e *Env) IsMac() bool {
	return e.OS == OSMac
}

func (e *Env) IsWin() bool {
	return e.OS == OSWindows
}

func (e *Env) IsLinux() bool {
	return e.OS == OSLinux
}

func (e *Env) IsUnknown() bool {
	return e.OS == OSUnknown
}
