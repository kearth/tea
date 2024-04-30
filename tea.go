/**
* Tea Framework
* Version 1.0.1
*
* Copyright 2018 - 2023, Kearth
* Golang 1.21
 */
package tea

// NewHTTPServer 获取HTTPServer实例
// panic 获取失败时抛出异常
func NewHTTPServer() *HTTPServer {
	var httpServer IContainer
	var err error
	// 获取HTTPServer实例
	if httpServer, err = IOC().Get(new(HTTPServer).Name()); err != nil {
		panic(err)
	}
	return httpServer.(*HTTPServer)
}
