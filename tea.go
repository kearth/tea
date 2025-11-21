package tea

import (
	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/internal/tea"
)

var (
	version = "0.0.9"
)

// Version 获取版本号
func Version() string {
	return version
}

// Drink 喝茶
func Drink(ctx kctx.Context, load func(ctx kctx.Context)) {
	// 启动框架
	tea.New(version, load).Run(ctx)
}
