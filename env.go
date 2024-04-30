package tea

const (
	// 正式环境
	EnvOnline  = "ONLINE"  // 线上
	EnvPreview = "PREVIEW" // 预览

	// 测试环境
	EnvTest = "TEST"
	EnvDev  = "DEV"

	// 系统环境
	SystemMac   = "mac"   // mac
	SystemWin   = "win"   // windows
	SystemLinux = "linux" // linux
	SystemOther = "other" // 其它
)

// DefaultEnvInfo 默认环境信息
var ei = &envInfo{
	env:    EnvDev,
	system: SystemLinux,
}

type envInfo struct {
	env    string
	system string
}

// 获取环境
func Env() string {
	return ei.env
}

// 设置环境
func SetEnv(env string) IError {
	m := map[string]struct{}{EnvOnline: {}, EnvPreview: {}, EnvTest: {}, EnvDev: {}}
	if _, ok := m[env]; !ok {
		return NotAllowType
	}
	ei.env = env
	return nil
}

// 获取系统环境
func System() string {
	return ei.system
}

// 设置系统环境
func SetSystem(system string) IError {
	m := map[string]struct{}{SystemMac: {}, SystemWin: {}, SystemLinux: {}, SystemOther: {}}
	if _, ok := m[system]; !ok {
		return NotAllowType
	}
	ei.system = system
	return nil
}
