package utils

import "fmt"

// SPF 格式化字符串
func SPF(f string, s ...any) string {
	return fmt.Sprintf(f, s...)
}

// Condition 三元运算
func Condition[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfThen 条件成立执行
func IfThen(condition bool, f func()) {
	if condition {
		f()
	}
}

// As 类型转换
func As[T any](v any) (T, bool) {
	if vt, ok := v.(T); ok {
		return vt, true
	}
	return *new(T), false
}

// AsPtr 类型转换为指针
func AsPtr[T any](v any) (*T, bool) {
	if vt, ok := v.(*T); ok {
		return vt, true
	}
	return nil, false
}

// IsThen 类型转换并执行
func IsThen[T any](v any, f func(T)) {
	if vt, ok := v.(T); ok {
		f(vt)
	}
}
