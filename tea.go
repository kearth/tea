/**
* Tea Framework
* Version 1.0.1
*
* Copyright 2018 - 2023, Kearth
* Golang 1.20
 */
package tea

import "github.com/kearth/tea/app"

// App Type
const (
	AppTypeHTTPServer  = "HTTPServer"
	AppTypeHTTPSServer = "HTTPSServer"
	AppTypeScript      = "Script"
	AppTypeCronTask    = "CronTask"
	AppTypeTestServer  = "TestServer"
)

// New
func New(tp string) app.App {
	switch tp {
	/*	case AppTypeCronTask:
			return
		case AppTypeHTTPServer:
			return new(app.HTTPServer)
		case AppTypeScript:
			return
		case AppTypeTestServer:
			return
	*/
	default:
		return new(app.HTTPServer)
	}
}
