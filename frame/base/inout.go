package base

// OutputFormat 默认输出格式
var out OutputFormat = &Output{}

// OutputFormat 输出格式接口
type OutputFormat interface {
	Output() any
}

// Output 输出结构体
type Output struct {
	Code int    `json:"code" dc:"错误码" example:"0"`  // 错误码
	Msg  string `json:"msg" dc:"错误信息" example:"Ok"` // 错误信息
	Data any    `json:"data" dc:"输出数据"`             // 输出数据
}

func (o *Output) Output() any {
	return o
}
