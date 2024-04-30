package main

import (
	"example/local/bootstrap"
	"example/local/router"

	"github.com/kearth/tea"
)

// 框架入口
func main() {
	// HTTP Server
	httpserver := tea.NewHTTPServer()
	httpserver.SetBootstrap(bootstrap.Bootstrap)
	httpserver.SetRouter(router.Router)
	httpserver.Start()
}
