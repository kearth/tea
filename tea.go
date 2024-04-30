/**
* Tea Framework
* Version 1.0.1
*
* Copyright 2018 - 2023, Kearth
* Golang 1.21
 */
package tea

import "github.com/kearth/tea/core"

// NewHTTPServer 获取HTTPServer实例
// panic 获取失败时抛出异常
func NewHTTPServer() *HTTPServer {
	var httpServer core.IContainer
	var err error
	// 获取HTTPServer实例
	if httpServer, err = core.IOC().Get(new(HTTPServer).Name()); err != nil {
		panic(err)
	}
	return httpServer.(*HTTPServer)
}
