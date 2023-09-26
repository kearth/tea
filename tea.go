/**
* Tea Framework
* Version 1.0.1
*
* Copyright 2018 - 2023, Kearth
* Golang 1.20
 */
package tea

import "github.com/kearth/tea/core"

// NewHTTPServer
func NewHTTPServer() *HTTPServer {
	var httpServer core.IContainer
	var err error
	if httpServer, err = core.IOC().Get("HTTPServer"); err != nil {
		panic(err)
	}
	return httpServer.(*HTTPServer)
}
