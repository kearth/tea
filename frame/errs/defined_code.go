package errs

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	// Defined errors.
	Success       = gerror.NewCode(gcode.CodeOK, "Success")
	ParamsEmpty   = gerror.NewCode(gcode.CodeMissingParameter, "Params Empty")     // 参数为空
	FormatError   = gerror.NewCode(gcode.CodeInvalidParameter, "Format Error")     // 格式错误
	UnknownError  = gerror.NewCode(gcode.CodeUnknown, "Unknown Error")             // 未知错误
	InternalError = gerror.NewCode(gcode.CodeInternalError, "Internal Error")      // 内部错误
	ConfigError   = gerror.NewCode(gcode.CodeInvalidConfiguration, "Config Error") // 配置错误
	NotFound      = gerror.NewCode(gcode.CodeNotFound, "Not Found")                // 未找到

	// 100 Framework errors.
	ConfigFileNotExists   = gerror.NewCode(gcode.New(100000, "", nil), "Config file not exists")   // 配置文件不存在
	DefaultServerNotFound = gerror.NewCode(gcode.New(100001, "", nil), "Default server not found") // 默认服务器未找到
	RouterNotFound        = gerror.NewCode(gcode.New(100002, "", nil), "Router not found")         // 路由器未找到
	ListenerNotFound      = gerror.NewCode(gcode.New(100003, "", nil), "Listener not found")       // 监听器未找到
	ResourceNotFound      = gerror.NewCode(gcode.New(100004, "", nil), "Resource not found")       // 资源未找到

	// RPCError          = &Error{100006, "rpc error"}               // 远程调用错误
	// NameRegistered    = &Error{200000, "the name has registered"} // 容器已注册
	// NameNotRegistered = &Error{200001, "the name not regisered"}  // 容器未注册
	// NotAllowType      = &Error{200002, "not allow type"}          // 不允许的类型
	// NoParserFound     = &Error{200004, "no parser found"}         // 解析器不存在
	// NoFileExt         = &Error{200005, "not file extension"}      // 文件后缀名错误
	// FrameworkError    = &Error{200006, "framework error"}         // 框架核心错误
)

// Wrap 包装错误
func Wrap(err error, message string) error {
	return gerror.Wrap(err, message)
}
