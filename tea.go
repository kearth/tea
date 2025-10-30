package tea

import (
	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/internal/tea"
)

var (
	version = "0.0.2"
)

// Version 获取版本号
func Version() string {
	return version
}

// Drink 喝茶
func Drink() {
	// 启动框架
	tea.New(kctx.New()).Run()
}
