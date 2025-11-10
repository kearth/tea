package main

import (
	"example/local/app/router"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea"
)

/*******************************
 * 框架入口
 * @author: kearth
 * 这是一个示例
 *******************************/
func main() {

	// 创建上下文
	ctx := kctx.New()

	// 启动框架
	tea.Drink(ctx, func() {
		// 启动服务器
		router.LoadRouter(ctx)
	})
}
