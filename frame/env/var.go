package env

import "github.com/gogf/gf/v2/util/gconv"

type Var struct {
	v any
}

func (v *Var) IsNil() bool {
	return v.v == nil
}

func (v *Var) String() string {
	return gconv.String(v.v)
}

func (v *Var) Int() int {
	return gconv.Int(v.v)
}

func (v *Var) Int64() int64 {
	return gconv.Int64(v.v)
}

func (v *Var) Float32() float32 {
	return gconv.Float32(v.v)
}

func (v *Var) Float64() float64 {
	return gconv.Float64(v.v)
}

func (v *Var) Bool() bool {
	return gconv.Bool(v.v)
}

func (v *Var) Any() any {
	return v.v
}

func (v *Var) Map() map[string]any {
	return gconv.Map(v.v)
}

func (v *Var) MapStrStr() map[string]string {
	return gconv.MapStrStr(v.v)
}
