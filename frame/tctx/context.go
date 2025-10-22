package tctx

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/utils"
)

var _ Context = (*tCtx)(nil)

type (
	// 上下文键
	CtxKey string

	// 上下文接口
	Context interface {
		context.Context
		Get(key string) string
		Set(key string, val string)
		Values() map[string]string
		Ctx() context.Context
		SetCtx(ctx context.Context)
		GetRequestID() string
		GetResponseID() string
	}

	// 上下文值
	tVal struct {
		RequestID  string
		ResponseID string
		AddTags    map[string]string
	}

	// 上下文实现
	tCtx struct {
		c context.Context
		m *tVal
	}
)

// 转换为gctx.StrKey
func (c CtxKey) GctxKey() gctx.StrKey {
	return gctx.StrKey(c)
}

const (
	RequestID  CtxKey = "RequestID"  // 请求ID
	ResponseID CtxKey = "ResponseID" // 响应ID
	AddTags    CtxKey = "AddTags"    // 附加参数
)

var (
	// 内部键
	innerKeys = []CtxKey{
		RequestID,
		ResponseID,
		AddTags,
	}

	DefaultRequestID  = base.DefaultID // 默认请求ID
	DefaultResponseID = base.DefaultID // 默认响应ID
)

// 创建上下文
func New() Context {
	v := &tVal{
		RequestID:  DefaultRequestID,
		ResponseID: DefaultResponseID,
		AddTags:    make(map[string]string),
	}
	return &tCtx{c: context.Background(), m: v}
}

// 创建上下文
func NewWithCtx(ctx context.Context) Context {
	v := &tVal{
		RequestID:  DefaultRequestID,
		ResponseID: DefaultResponseID,
		AddTags:    make(map[string]string),
	}
	return (&tCtx{c: ctx, m: v}).getInnerKeyValueToMap()
}

// 获取内部上下文的值
func (t *tCtx) getInnerKeyValueToMap() *tCtx {
	for _, key := range innerKeys {
		if v := t.c.Value(key); v != nil {
			if v == AddTags {
				vm, _ := utils.As[map[string]string](v)
				for k, vs := range vm {
					t.Set(k, vs)
				}
			} else {
				vs, _ := utils.As[string](v)
				t.Set(string(key), vs)
			}
		}
	}
	return t
}

// 设置上下文
func (t *tCtx) SetCtx(ctx context.Context) {
	t.c = ctx
}

// WithCancel 创建一个上下文，当调用 cancel 时，上下文会被取消
func WithCancel(parent Context) (Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent.Ctx())
	parent.SetCtx(ctx)
	return parent, cancel
}

// 创建一个上下文，当超时后，上下文会被取消
func WithDeadline(parent Context, deadline time.Time) (Context, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(parent.Ctx(), deadline)
	parent.SetCtx(ctx)
	return parent, cancel
}

// 创建一个上下文，当超时后，上下文会被取消
func WithTimeout(parent Context, timeout time.Duration) (Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent.Ctx(), timeout)
	parent.SetCtx(ctx)
	return parent, cancel
}

// 创建一个上下文，当父上下文取消时，子上下文也会被取消
func WithValue(parent Context, key any, val any) Context {
	ctx := context.WithValue(parent.Ctx(), key, val)
	parent.SetCtx(ctx)
	return parent
}

// Done 返回一个通道，当上下文结束时，该通道会被关闭
func (t *tCtx) Done() <-chan struct{} {
	return t.c.Done()
}

// Err 返回上下文结束时的错误
func (t *tCtx) Err() error {
	return t.c.Err()
}

// 获取上下文中的值
func (t *tCtx) Value(key any) any {
	return t.c.Value(key)
}

// 取消上下文
func (t *tCtx) Deadline() (deadline time.Time, ok bool) {
	return t.c.Deadline()
}

// 获取上下文中的值
func (t *tCtx) Get(key string) string {
	// 不存在则返回空字符串
	return t.m.AddTags[key]
}

// 设置上下文中的值
func (t *tCtx) Set(key string, val string) {
	// 存在则覆盖
	t.m.AddTags[key] = val
}

// 获取上下文中的值
func (t *tCtx) Values() map[string]string {
	return t.m.AddTags
}

// 获取上下文
func (t *tCtx) Ctx() context.Context {
	ctx := t.c
	for _, key := range innerKeys {
		switch key {
		case RequestID:
			ctx = context.WithValue(ctx, key, t.m.RequestID)
		case ResponseID:
			ctx = context.WithValue(ctx, key, t.m.ResponseID)
		case AddTags:
			ctx = context.WithValue(ctx, key, t.m.AddTags)
		}
	}
	t.SetCtx(ctx)
	return ctx
}

func (t *tCtx) GetRequestID() string {
	return t.m.RequestID
}

func (t *tCtx) GetResponseID() string {
	return t.m.ResponseID
}
