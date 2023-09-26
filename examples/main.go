package main

import (
	"example/local/bootstrap"
	"example/local/router"

	"github.com/kearth/tea"
)

func main() {
	// HTTP Server
	httpserver := tea.NewHTTPServer()
	httpserver.SetBootstrap(bootstrap.Bootstrap)
	httpserver.SetRouter(router.Router)
	httpserver.Start()
}
