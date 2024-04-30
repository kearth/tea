/**
* Tea Framework
* Version 1.0.1
*
* Copyright 2018 - 2023, Kearth
* Golang 1.21
 */
package tea

// HTTPServer 获取HTTPServer实例
// panic 获取失败时抛出异常
func HTTPServer() *httpServer {
	var hs IContainer
	var err error
	// 获取HTTPServer实例
	if hs, err = IOC().Get(new(httpServer).Name()); err != nil {
		panic(err)
	}
	return hs.(*httpServer)
}
