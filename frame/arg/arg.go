package arg

import "reflect"

// Arg 参数
type Arg struct {
	Val  any
	Type reflect.Type
}

func NewArg(val any) *Arg {
	return &Arg{
		Val:  val,
		Type: reflect.TypeOf(val),
	}
}
