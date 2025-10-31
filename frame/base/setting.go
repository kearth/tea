package base

import "github.com/kearth/klib/kerr"

var (
	configPath = "./conf/tea.toml" // 配置文件路径
)

// SetConfigPath 设置配置文件路径
func SetConfigPath(path string) {
	configPath = path
}

// GetConfigPath 获取配置文件路径
func GetConfigPath() string {
	return configPath
}

const (
	DefaultServer = "default" // 默认服务器
)

const (
	DefaultID     = "-1"     // 默认ID
	EnvModeNormal = "normal" // 正常模式
	EnvModeDebug  = "debug"  // 调试模式

	OSMac       = "darwin"    // macOS
	OSLinux     = "linux"     // Linux
	OSWindows   = "windows"   // Windows
	OSAndroid   = "android"   // Android
	OSFreeBSD   = "freebsd"   // FreeBSD
	OSDragonfly = "dragonfly" // Dragonfly
	OSNetBSD    = "netbsd"    // NetBSD
	OSPlain     = "plan9"     // Plan 9
	OSSolaris   = "solaris"   // Solaris
	OSOpenBSD   = "openbsd"   // OpenBSD
	OSUnknown   = "unknown"   // 未知系统

)

// Base error definitions
var (
	ConfigFileNotExists   = kerr.New(600001, "Config file not exists").WithDisplay("配置文件不存在")
	DefaultServerNotFound = kerr.New(600002, "Default server not found").WithDisplay("默认服务器未找到")
	RouterNotFound        = kerr.New(600003, "Router not found").WithDisplay("路由器未找到")
	ListenerNotFound      = kerr.New(600004, "Listener not found").WithDisplay("监听器未找到")
	ResourceNotFound      = kerr.New(600005, "Resource not found").WithDisplay("资源未找到")
	StepError             = kerr.New(600006, "Step error").WithDisplay("步骤错误")
)
