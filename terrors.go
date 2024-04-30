package tea

import (
	"errors"
	"fmt"
)

var (
	// 0 代表成功,除0以外的其他值都代表失败
	Succ = &Error{0, "success"}
	// 错误码6位, 100000-999999
	// 100 开头是预定义错误
	ParamsEmpty = &Error{100000, "params empty"} // 参数为空
	ParamsError = &Error{100001, "params error"} // 参数错误
	FormatError = &Error{100002, "format error"} // 参数格式错误
	AssertError = &Error{100003, "assert error"} // 类型断言错误
	TransError  = &Error{100004, "trans error"}  // 类型转换错误
	SystemError = &Error{100005, "system error"} // 系统错误
	RPCError    = &Error{100006, "rpc error"}    // 远程调用错误
	ConfError   = &Error{100007, "conf error"}   // 内部配置错误
	NoData      = &Error{100008, "no data"}      // 没有数据
	// 200 开头是框架core错误
	NameRegistered    = &Error{200000, "the name has registered"} // 容器已注册
	NameNotRegistered = &Error{200001, "the name not regisered"}  // 容器未注册
	NotAllowType      = &Error{200002, "not allow type"}          // 不允许的类型
	ConfNotExists     = &Error{200003, "conf not exists"}         // 配置文件不存在
	NoParserFound     = &Error{200004, "no parser found"}         // 解析器不存在
	NoFileExt         = &Error{200005, "not file extension"}      // 文件后缀名错误
	FrameworkError    = &Error{200006, "framework error"}         // 框架核心错误
	// 300 开头是框架其他部分错误
	// 400 开头是用户自定义错误
)

// IError error interface
type IError interface {
	error
	Code() int             // 错误码
	Wrap(err error) IError // 包装错误
	Unwrap() error         // 获取原始错误
}

// check 错误检查
var _ IError = new(Error)

// Error
type Error struct {
	code int
	msg  string
}

// NewErr 创建错误
// 自定义错误码必须大于等于400000，小于等于999999
func NewErr(c int, m string) *Error {
	if c < 400000 || c > 999998 {
		// 999999 代表错误码超限
		c = 999999
	}
	return &Error{code: c, msg: m}
}

// Error 错误信息
func (e *Error) Error() string {
	return e.msg
}

// Code 错误码
func (e *Error) Code() int {
	return e.code
}

// Wrap 包装错误
func (e *Error) Wrap(err error) IError {
	if err != nil {
		e.msg = fmt.Sprintf("%s:%s", e.msg, err)
	}
	return e
}

// Unwrap 获取原始错误
func (e *Error) Unwrap() error {
	return errors.New(e.Error())
}
