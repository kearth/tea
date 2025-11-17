package base

import (
	"github.com/gogf/gf/v2/encoding/gjson"
)

// OutputFormat 默认输出格式
var out OutputFormat = &Output{}

// OutputFormat 输出格式接口
type OutputFormat interface {
	SetCode(code int)
	SetMsg(msg string)
	SetData(data any)
	String() string
}

func Response() OutputFormat {
	return out
}

// Output 输出结构体
type Output struct {
	Code int    `json:"code" dc:"错误码" example:"0"`  // 错误码
	Msg  string `json:"msg" dc:"错误信息" example:"OK"` // 错误信息
	Data any    `json:"data" dc:"输出数据"`             // 输出数据
}

func (o *Output) SetCode(code int) {
	o.Code = code
}

func (o *Output) SetMsg(msg string) {
	o.Msg = msg
}

func (o *Output) SetData(data any) {
	o.Data = data
}

func (o *Output) String() string {
	s, _ := gjson.EncodeString(o)
	return s
}
