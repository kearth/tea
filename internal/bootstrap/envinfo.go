package bootstrap

import (
	"bytes"
	"context"
	"io"
	"os"
	"runtime"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/t"
	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/frame/utils"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	// 组件接口
	_ container.Component = (*EnvInfo)(nil)

	// 实例
	envInfoInstance = &EnvInfo{}
)

// 环境
type EnvInfo struct {
	container.BaseObject
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

func LoadEnvInfo() *EnvInfo {
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
func (e *EnvInfo) Init(ctx tctx.Context) error {
	e.SetName("EnvInfo")
	// 解析配置文件
	cfg, err := ParseTOML(base.ConfigPath)
	if err != nil {
		return err
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
		e.Address = utils.SPF("%s:%d", e.IP, e.Port)
	} else {
		e.Address = utils.SPF("%s:%d", e.Host, e.Port)
	}

	envInfoInstance.SystemVersion = e.getSystemVersion(ctx)
	envInfoInstance.OS = runtime.GOOS
	envInfoInstance.Arch = runtime.GOARCH
	envInfoInstance.CPU = runtime.NumCPU()
	envInfoInstance.Hostname, _ = os.Hostname()
	envInfoInstance.PID = os.Getpid()

	// 设置默认服务器
	defaultServer := cfg.MustGet(ctx, "framework.default_server", base.DefaultServer).String()
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

// 设置debug模式
func (e *EnvInfo) SetDebug() {
	e.Mode = base.EnvModeDebug
}

// 设置RootDir
func (e *EnvInfo) SetRootDir(path string) bool {
	if path == "" {
		return false
	}
	e.RootDir = path
	return true
}

// 设置ResourcesDir
func (e *EnvInfo) SetResourcesDir(path string) bool {
	if path == "" {
		return false
	}
	e.ResourcesDir = path
	return true
}
