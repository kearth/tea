/**
* Tea Framework
* Version 1.0.1
*
* Copyright 2018 - 2023, Kearth
* Golang 1.20
 */
package tea

import (
	"github.com/kearth/tea/app"
	"github.com/kearth/tea/tserver"
)

type AppType string

var (
	_ app.IApp = new(tserver.HTTPServer)
)

const (
	AppHttpServer AppType = "http_server"
)

// New
func NewServer(at AppType) app.IServer {
	var is app.IServer
	switch at {
	case AppHttpServer:
		is = new(tserver.HTTPServer)
	default:
		is = new(tserver.HTTPServer)
	}
	return is
}
