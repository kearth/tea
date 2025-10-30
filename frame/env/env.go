package env

import (
	"sync"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/internal/bootstrap"
)

var envInfo *bootstrap.EnvInfo
var once sync.Once

// loadEnvInfo 用于懒加载环境信息
func loadEnvInfo() *bootstrap.EnvInfo {
	once.Do(func() {
		envInfo = bootstrap.LoadEnvInfo()
	})
	return envInfo
}

// IsDebug 是否调试模式
func IsDebug() bool {
	return loadEnvInfo().Mode == base.EnvModeDebug
}

// IsNormal 是否正常模式
func IsNormal() bool {
	return !IsDebug()
}

// IsRelease 是否正式模式
func IsRelease() bool {
	return loadEnvInfo().Mode == base.EnvModeNormal
}

// IsMac 是否Mac系统
func IsMac() bool {
	return loadEnvInfo().OS == base.OSMac
}

// IsWin 是否Windows系统
func IsWin() bool {
	return loadEnvInfo().OS == base.OSWindows
}

// IsLinux 是否Linux系统
func IsLinux() bool {
	return loadEnvInfo().OS == base.OSLinux
}

// IsUnknown 是否未知系统
func IsUnknown() bool {
	return loadEnvInfo().OS == base.OSUnknown
}

// GetPort 获取端口
func Port() int {
	return loadEnvInfo().Port
}

// GetIP 获取IP地址
func GetIP() string {
	return loadEnvInfo().IP
}

// GetHost 获取主机名
func GetHost() string {
	return loadEnvInfo().Host
}

// GetServerName 获取服务器名称
func GetServerName() string {
	return loadEnvInfo().ServerName
}

// GetServerVersion 获取服务器版本
func GetServerVersion() string {
	return loadEnvInfo().ServerVersion
}

// GetServerMode 获取服务器模式
func GetServerMode() string {
	return loadEnvInfo().Mode
}

// GetRootDir 获取根目录
func GetRootDir() string {
	return loadEnvInfo().RootDir
}

// GetResourcesDir 获取资源目录
func GetResourcesDir() string {
	return loadEnvInfo().ResourcesDir
}

// GetCfg 获取配置信息
func Cfg() *gcfg.Config {
	return loadEnvInfo().Cfg
}
