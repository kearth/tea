package tea

import (
	"time"
)

/*******************************************
 * 最小的执行单元
 * 有名称，计时统计，有执行方法
 * 一个单元不应该并发执行
 *******************************************/

var _ IUnit = (*unit)(nil)

// IUnit 单元接口
type IUnit interface {
	Cost() time.Duration
	Name() string
	Call(ctx Context, input ...any) (any, Error)
}

// Unit 单元
type unit struct {
	name  string                                                  // 名称
	start time.Time                                               // 开始时间
	end   time.Time                                               // 结束时间
	cost  time.Duration                                           // 耗时
	err   IError                                                  // 错误
	ctx   Context                                                 // 上下文
	fn    func(ctx Context, input ...any) (output any, err Error) // 执行方法
}

// Cost 计算耗时
func (u *unit) Cost() time.Duration {
	// 计算耗时
	if u.start.IsZero() {
		u.start = time.Now()
	}
	if u.end.IsZero() {
		u.end = time.Now()
	}
	// 计算耗时
	if u.cost == 0 {
		u.cost = u.end.Sub(u.start)
	}
	return u.cost
}

// Name 获取名称
func (u *unit) Name() string {
	return u.name
}

// Call 执行方法
func (u *unit) Call(ctx Context, input ...any) (any, Error) {
	u.start = time.Now()
	// 执行方法
	output, err := u.fn(ctx, input)
	if err != *Succ {
		u.err = &err
	}
	u.end = time.Now()
	u.cost = u.end.Sub(u.start)
	return output, err
}

// NewUnit 创建单元
func NewUnit(ctx Context, name string, fn func(ctx Context, input ...any) (output any, err Error)) IUnit {
	if name == "" {
		name = "no name unit"
	}
	return &unit{
		name: name,
		fn:   fn,
		err:  Succ,
		ctx:  ctx,
	}
}
