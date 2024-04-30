package core

import (
	"context"
	"time"
)

// check
var _ Context = (*tCtx)(nil)

// 上下文接口
type Context interface {
	context.Context
	Get(key string) string
	Set(key string, val string)
	Values() map[string]string
	Ctx() context.Context
}

// 上下文实现
type tCtx struct {
	it context.Context
	m  map[string]string
}

// 创建上下文
func Background() Context {
	return &tCtx{it: context.Background(), m: make(map[string]string)}
}

// WithCancel 创建一个上下文，当调用 cancel 时，上下文会被取消
func WithCancel(parent Context) (Context, context.CancelFunc) {
	it, cancel := context.WithCancel(parent.Ctx())
	return &tCtx{it: it}, cancel
}

// 创建一个上下文，当超时后，上下文会被取消
func WithDeadline(parent Context, deadline time.Time) (Context, context.CancelFunc) {
	it, cancel := context.WithDeadline(parent.Ctx(), deadline)
	return &tCtx{it: it}, cancel
}

// 创建一个上下文，当超时后，上下文会被取消
func WithTimeout(parent Context, timeout time.Duration) (Context, context.CancelFunc) {
	it, cancel := context.WithTimeout(parent.Ctx(), timeout)
	return &tCtx{it: it}, cancel
}

// 创建一个上下文，当父上下文取消时，子上下文也会被取消
func WithValue(parent Context, key interface{}, val interface{}) Context {
	it := context.WithValue(parent.Ctx(), key, val)
	return &tCtx{it: it}
}

// Done 返回一个通道，当上下文结束时，该通道会被关闭
func (t *tCtx) Done() <-chan struct{} {
	return t.it.Done()
}

// Err 返回上下文结束时的错误
func (t *tCtx) Err() error {
	return t.it.Err()
}

// 获取上下文中的值
func (t *tCtx) Value(key interface{}) interface{} {
	return t.it.Value(key)
}

// 取消上下文
func (t *tCtx) Deadline() (deadline time.Time, ok bool) {
	return t.it.Deadline()
}

// 获取上下文中的值
func (t *tCtx) Get(key string) string {
	// 不存在则返回空字符串
	return t.m[key]
}

// 设置上下文中的值
func (t *tCtx) Set(key string, val string) {
	// 存在则覆盖
	t.m[key] = val
}

// 获取上下文中的值
func (t *tCtx) Values() map[string]string {
	return t.m
}

// 获取上下文
func (t *tCtx) Ctx() context.Context {
	return t.it
}
