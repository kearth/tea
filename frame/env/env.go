package env

import (
	"sync"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/kearth/tea/frame/base"
)

// 环境
type Env struct {
	Version       string
	ServerName    string `json:"server_name"`
	ServerVersion string `json:"server_version"`
	Mode          string `json:"mode"`
	OS            string
	Port          int    `json:"port"`
	IP            string `json:"ip"`
	Address       string
	Arch          string
	SystemVersion string
	CPU           int
	PID           int
	Hostname      string
	RootDir       string `json:"root_dir"`
	ResourcesDir  string `json:"resources_dir"`
	ServerType    string `json:"server_type"`
	Cfg           *gcfg.Config
}

var instance *Env
var once sync.Once

// Init
func Init(e *Env) {
	instance = e
}

// Instance 用于懒加载环境信息
func Instance() *Env {
	return instance
}

// IsDebug 是否调试模式
func IsDebug() bool {
	return Instance().Mode == base.EnvModeDebug
}

// IsNormal 是否正常模式
func IsNormal() bool {
	return !IsDebug()
}

// IsRelease 是否正式模式
func IsRelease() bool {
	return Instance().Mode == base.EnvModeNormal
}

// IsMac 是否Mac系统
func IsMac() bool {
	return Instance().OS == base.OSMac
}

// IsWin 是否Windows系统
func IsWin() bool {
	return Instance().OS == base.OSWindows
}

// IsLinux 是否Linux系统
func IsLinux() bool {
	return Instance().OS == base.OSLinux
}

// IsUnknown 是否未知系统
func IsUnknown() bool {
	return Instance().OS == base.OSUnknown
}

// GetOS 获取操作系统
func GetOS() string {
	return Instance().OS
}

// GetPort 获取端口
func Port() int {
	return Instance().Port
}

// GetIP 获取IP地址
func GetIP() string {
	return Instance().IP
}

// GetServerName 获取服务器名称
func GetServerName() string {
	return Instance().ServerName
}

// GetServerVersion 获取服务器版本
func GetServerVersion() string {
	return Instance().ServerVersion
}

// GetServerMode 获取服务器模式
func GetServerMode() string {
	return Instance().Mode
}

// GetRootDir 获取根目录
func GetRootDir() string {
	return Instance().RootDir
}

// GetResourcesDir 获取资源目录
func GetResourcesDir() string {
	return Instance().ResourcesDir
}

// GetCfg 获取配置信息
func Cfg() *gcfg.Config {
	return Instance().Cfg
}

// GetServerType 获取服务器类型
func GetServerType() string {
	return Instance().ServerType
}
