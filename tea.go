package tea

import (
	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/internal/tea"
)

var (
	version = "0.0.5"
)

// Version 获取版本号
func Version() string {
	return version
}

// Drink 喝茶
func Drink() {
	// 启动框架
	tea.New(version).Run(kctx.New())
}
