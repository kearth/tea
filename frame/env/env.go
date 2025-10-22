package env

import (
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/internal/bootstrap"
)

// IsDebug 是否调试模式
func IsDebug() bool {
	return bootstrap.LoadEnvInfo().Mode == base.EnvModeDebug
}

// IsNormal 是否正常模式
func IsNormal() bool {
	return !IsDebug()
}

// IsRelease 是否正式模式
func IsMac() bool {
	return bootstrap.LoadEnvInfo().OS == base.OSMac
}

// IsWin 是否Windows系统
func IsWin() bool {
	return bootstrap.LoadEnvInfo().OS == base.OSWindows
}

// IsLinux 是否Linux系统
func IsLinux() bool {
	return bootstrap.LoadEnvInfo().OS == base.OSLinux
}

// IsUnknown 是否未知系统
func IsUnknown() bool {
	return bootstrap.LoadEnvInfo().OS == base.OSUnknown
}

// Port 端口
func Port() int {
	return bootstrap.LoadEnvInfo().Port
}

// IP IP地址
func IP() string {
	return bootstrap.LoadEnvInfo().IP
}

// Host 主机名
func Host() string {
	return bootstrap.LoadEnvInfo().Host
}

// ServerName 服务器名称
func ServerName() string {
	return bootstrap.LoadEnvInfo().ServerName
}

// ServerVersion 服务器版本
func ServerVersion() string {
	return bootstrap.LoadEnvInfo().ServerVersion
}

// ServerMode 服务器模式
func ServerMode() string {
	return bootstrap.LoadEnvInfo().Mode
}

// RootDir 根目录
func RootDir() string {
	return bootstrap.LoadEnvInfo().RootDir
}

// ResourcesDir 资源目录
func ResourcesDir() string {
	return bootstrap.LoadEnvInfo().ResourcesDir
}

// Cfg 配置信息
func Cfg() *gcfg.Config {
	return bootstrap.LoadEnvInfo().Cfg
}
