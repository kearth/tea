package tea

import (
	"runtime"

	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/internal/tea"
)

// Drink 喝茶
func Drink() {

	// 设置最大CPU核心数
	_ = runtime.GOMAXPROCS(runtime.NumCPU())

	// 启动框架
	tea.New().Run(tctx.New())
}
