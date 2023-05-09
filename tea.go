/**
* Tea Framework
* Version 1.0.1
*
* Copyright 2018 - 2023, Kearth
* Golang 1.20
 */
package tea

// NewHTTPServer
func NewHTTPServer() *HTTPServer {
	var httpServer IContainer
	var err error
	if httpServer, err = IOC().Get("HTTPServer"); err != nil {
		panic(err)
	}
	return httpServer.(*HTTPServer)
}
