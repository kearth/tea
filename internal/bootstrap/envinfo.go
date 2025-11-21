package bootstrap

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

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
		Unit:      kunit.NewUnit(EnvName).SetRole(kunit.RoleComponent),
		LogConfig: env.LogConfig{},
		Config:    env.Config{},
		Env: env.Env{
			Var: map[string]any{},
		},
	}
)

// 环境
type EnvInfo struct {
	kunit.Unit
	env.Env
	env.LogConfig
	env.Config
}

// Env 获取实例
func Env() *EnvInfo {
	return envInfoInstance
}

// ParseTOML 解析配置文件
func (e *EnvInfo) ParseTOML(name string) (*gcfg.Config, kerr.Error) {
	if name == "" {
		return nil, base.ConfigFileNotExists.Wrap(fmt.Errorf("config file name is empty"))
	}

	var path string
	var err kerr.Error
	if path, err = e.IsValidFile(name); err != nil {
		return nil, err
	}
	g.Cfg(name).GetAdapter().(*gcfg.AdapterFile).SetFileName(path)
	return g.Cfg(name), nil
}

func (e *EnvInfo) SetVersion(version string) *EnvInfo {
	e.Version = version
	return e
}

func (e *EnvInfo) Path(name string) string {
	return gfile.Join(base.GetConfigPath(), name)
}

func (e *EnvInfo) IsExist(name string) bool {
	return gfile.Exists(e.Path(name))
}

func (e *EnvInfo) IsToml(name string) bool {
	return gfile.ExtName(name) == "toml"
}

func (e *EnvInfo) IsValidFile(name string) (string, kerr.Error) {
	if !e.IsExist(name) {
		return "", base.ConfigFileNotExists.Wrap(fmt.Errorf("config file not exists, path: %s", name))
	}
	if !e.IsToml(name) {
		return "", base.ConfigFileNotExists.Wrap(fmt.Errorf("config file must be toml format, path: %s", name))
	}
	return e.Path(name), nil
}

// 初始化环境
func (e *EnvInfo) Setup(ctx kctx.Context) kerr.Error {
	// 解析配置文件
	cfg, err := e.ParseTOML(base.ConfigName)
	if err != nil {
		return err
	}

	// 读取配置文件
	e.Config.Server = cfg.MustGet(ctx, "config.server").String()
	e.Config.Log = cfg.MustGet(ctx, "config.log").String()
	e.Config.Var = cfg.MustGet(ctx, "config.var").String()
	e.Cfg = cfg

	var serverCfg *gcfg.Config
	if e.Config.Server != "" {
		serverCfg, err = e.ParseTOML(e.Config.Server)
		if err != nil {
			return err
		}
	}
	// 后赋值，避免引用的对象被覆盖
	if serverCfg == nil {
		serverCfg = cfg
	}
	if !serverCfg.Available(ctx, "server", "name") {
		return base.ConfigKeyNotExists.Wrap(fmt.Errorf("server.name"))
	}
	e.ServerName = serverCfg.MustGet(ctx, "server.name", "tea").String()
	e.ServerVersion = serverCfg.MustGet(ctx, "server.version", "1.0.0").String()
	e.RootDir = serverCfg.MustGet(ctx, "server.root_dir", gfile.Pwd()).String()
	if e.RootDir == "." || e.RootDir == "" || e.RootDir == "/" {
		e.RootDir = gfile.Pwd()
	}
	e.ResourcesDir = serverCfg.MustGet(ctx, "server.resources_dir", "./resources").String()
	e.Mode = serverCfg.MustGet(ctx, "server.mode", base.EnvModeNormal).String()
	// 调试模式
	if e.Mode != base.EnvModeDebug {
		e.Mode = base.EnvModeNormal
	}
	e.IP = serverCfg.MustGet(ctx, "server.ip", base.DefaultIP).String()
	e.Port = serverCfg.MustGet(ctx, "server.port", base.DefaultPort).Int()
	e.Address = fmt.Sprintf("%s:%d", e.IP, e.Port)
	e.ServerType = serverCfg.MustGet(ctx, "server.server_type", base.ServerTypeHTTP).String()
	e.SystemVersion = e.getSystemVersion(ctx)
	e.OS = runtime.GOOS
	e.Arch = runtime.GOARCH
	e.CPU = runtime.NumCPU()
	e.Hostname, _ = os.Hostname()
	e.PID = os.Getpid()

	var logCfg *gcfg.Config
	if e.Config.Log != "" {
		logCfg, err = e.ParseTOML(e.Config.Log)
		if err != nil {
			return err
		}
	}
	if logCfg == nil {
		logCfg = cfg
	}
	if !logCfg.Available(ctx, "log", "path") {
		return base.ConfigKeyNotExists.Wrap(fmt.Errorf("log.path"))
	}

	// log 配置
	e.LogConfig.Path = logCfg.MustGet(ctx, "log.path", "./log").String()
	e.LogConfig.File = logCfg.MustGet(ctx, "log.file", "{Y-m-d}.log").String()
	e.LogConfig.Prefix = logCfg.MustGet(ctx, "log.prefix", "tea").String()
	e.LogConfig.Level = logCfg.MustGet(ctx, "log.level", "all").String()
	e.LogConfig.TimeFormat = logCfg.MustGet(ctx, "log.time_format", time.DateTime).String()
	e.LogConfig.CtxKeys = logCfg.MustGet(ctx, "log.ctx_keys", []string{}).Strings()
	e.LogConfig.Header = logCfg.MustGet(ctx, "log.header", true).Bool()
	e.LogConfig.Stdout = logCfg.MustGet(ctx, "log.stdout", true).Bool()
	e.LogConfig.StdoutColorDisabled = logCfg.MustGet(ctx, "log.stdout_color_disabled", false).Bool()
	e.LogConfig.WriterColorEnable = logCfg.MustGet(ctx, "log.writer_color_enable", false).Bool()

	// 初始化日志配置
	logger := klog.Logger()
	logger.SetPath(e.LogConfig.Path)
	logger.SetFile(e.LogConfig.File)
	logger.SetTimeFormat(e.LogConfig.TimeFormat)
	logger.SetPrefix(e.LogConfig.Prefix)
	logger.SetLevelStr(e.LogConfig.Level)
	logger.SetCtxKeys(e.LogConfig.CtxKeys)
	logger.SetHeaderPrint(e.LogConfig.Header)
	logger.SetStdoutPrint(e.LogConfig.Stdout)
	logger.SetStdoutColorDisabled(e.LogConfig.StdoutColorDisabled)
	logger.SetWriterColorEnable(e.LogConfig.WriterColorEnable)

	// 解析变量配置文件
	var varCfg *gcfg.Config
	if e.Config.Var != "" {
		varCfg, err = e.ParseTOML(e.Config.Var)
		if err != nil {
			return err
		}
		e.Env.Var = varCfg.MustGet(ctx, "var").Map()
	}
	if varCfg == nil {
		varCfg = cfg
	}
	if varCfg.Available(ctx, "var") {
		e.Env.Var = varCfg.MustGet(ctx, "var").Map()
	}

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
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("LogPath:                  [ %s ]", e.LogConfig.Path))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("LogFile:                  [ %s ]", e.LogConfig.File))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("LogPrefix:                [ %s ]", e.LogConfig.Prefix))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("LogLevel:                 [ %s ]", e.LogConfig.Level))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("LogLevel:                 [ %s ]", e.LogConfig.Level))
	klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("ConfigPath:               [ %s ]", e.Path(base.ConfigName)))
	if e.Config.Server != "" {
		klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("ServerConfigPath:         [ %s ]", e.Path(e.Config.Server)))
	}
	if e.Config.Log != "" {
		klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("LogConfigPath:            [ %s ]", e.Path(e.Config.Log)))
	}
	if e.Config.Var != "" {
		klog.ColorPrint(ctx, klog.Cyan, fmt.Sprintf("VarConfigPath:            [ %s ]", e.Path(e.Config.Var)))
	}
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
