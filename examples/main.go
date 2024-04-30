package main

import (
	"example/local/bootstrap"
	"example/local/router"

	"github.com/kearth/tea"
)

/*******************************
 * 框架入口
 * @author: kearth
 * 这是一个示例
 *******************************/
func main() {
	// 一个 HTTP Server
	httpserver := tea.HTTPServer()
	httpserver.SetBootstrap(bootstrap.Bootstrap)
	httpserver.SetRouter(router.Router)
	httpserver.Start()
}
