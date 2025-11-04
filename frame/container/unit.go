package container

import (
	"time"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
)

/*******************************************
 * 最小的执行单元
 * 有名称，计时统计，有执行方法
 * 一个单元不应该并发执行
 *******************************************/

type Role string

const (
	RoleComponent Role = "Component"
	RoleModule    Role = "Module"
	RoleRouter    Role = "Router"
	RoleServer    Role = "Server"
	RoleUnit      Role = "Unit"
	RolePlugin    Role = "Plugin"
	RoleUnknown   Role = "Unknown"
)

type Fn func(ctx kctx.Context, input ...any) (output any, err kerr.Error)

// 修正后
type Unit interface {
	Cost() time.Duration
	Name() string
	Role() Role
	Call(ctx kctx.Context, input ...any) (any, kerr.Error)
	SetRole(role Role)
	SetName(name string)
	SetFn(fn Fn)
	Setup(ctx kctx.Context) kerr.Error
}

// Unit 单元
type unit struct {
	name  string        // 名称
	role  Role          // 角色
	start time.Time     // 开始时间
	end   time.Time     // 结束时间
	cost  time.Duration // 耗时
	err   kerr.Error    // 错误
	fn    Fn            // 执行方法
}

func (u *unit) Setup(ctx kctx.Context) kerr.Error {
	return nil
}

// Role 获取角色
func (u *unit) Role() Role {
	return u.role
}

// SetRole 设置角色
func (u *unit) SetRole(role Role) {
	u.role = role
}

// Cost 计算耗时（仅在执行后有效）
func (u *unit) Cost() time.Duration {
	// 未执行过的单元耗时为 0
	if u.start.IsZero() || u.end.IsZero() {
		return 0
	}
	// 已执行的单元返回预计算的耗时（避免重复计算）
	return u.cost
}

// Name 获取名称
func (u *unit) Name() string {
	return u.name
}

// SetName 设置名称
func (u *unit) SetName(name string) {
	u.name = name
}

// Call 执行单元
func (u *unit) Call(ctx kctx.Context, input ...any) (any, kerr.Error) {
	var output any
	var err kerr.Error
	u.start = time.Now()
	if u.fn != nil {
		output, err = u.fn(ctx, input)
		u.err = err // 直接存储 error（若 IError 是 interface{ Error() string }）
	} else {
		u.err = kerr.DependencyMissing
	}
	u.end = time.Now()
	u.cost = u.end.Sub(u.start)
	return output, err
}

// SetFn 设置执行方法
func (u *unit) SetFn(fn Fn) {
	u.fn = fn
}

// NewUnit 创建单元
func NewUnit(name string, fn ...Fn) Unit {
	if name == "" {
		name = "NoNameUnit"
	}

	u := &unit{
		name: name,
		err:  kerr.Succ,
		role: RoleUnit,
	}
	if len(fn) > 0 && fn[0] != nil {
		u.fn = fn[0]
	}
	return u
}
