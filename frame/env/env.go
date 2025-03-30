package env

import (
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/kearth/tea/frame/defined"
	"github.com/kearth/tea/internal/envinfo"
)

// IsDebug 是否调试模式
func IsDebug() bool {
	return envinfo.Instance().Mode == defined.EnvModeDebug
}

// IsNormal 是否正常模式
func IsNormal() bool {
	return !IsDebug()
}

// IsRelease 是否正式模式
func IsMac() bool {
	return envinfo.Instance().OS == defined.OSMac
}

// IsWin 是否Windows系统
func IsWin() bool {
	return envinfo.Instance().OS == defined.OSWindows
}

// IsLinux 是否Linux系统
func IsLinux() bool {
	return envinfo.Instance().OS == defined.OSLinux
}

// IsUnknown 是否未知系统
func IsUnknown() bool {
	return envinfo.Instance().OS == defined.OSUnknown
}

// Port 端口
func Port() int {
	return envinfo.Instance().Port
}

// IP IP地址
func IP() string {
	return envinfo.Instance().IP
}

// Host 主机名
func Host() string {
	return envinfo.Instance().Host
}

// ServerName 服务器名称
func ServerName() string {
	return envinfo.Instance().ServerName
}

// ServerVersion 服务器版本
func ServerVersion() string {
	return envinfo.Instance().ServerVersion
}

// ServerMode 服务器模式
func ServerMode() string {
	return envinfo.Instance().Mode
}

// RootDir 根目录
func RootDir() string {
	return envinfo.Instance().RootDir
}

// ResourcesDir 资源目录
func ResourcesDir() string {
	return envinfo.Instance().ResourcesDir
}

// Cfg 配置信息
func Cfg() *gcfg.Config {
	return envinfo.Instance().Cfg
}
