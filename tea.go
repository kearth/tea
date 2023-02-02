/**
* Tea Framework
* Version 1.0.1
*
* Copyright 2018 - 2023, Kearth
* Golang 1.20
 */
package tea

import "github.com/kearth/tea/app"

func New(tp string) app.App {
	switch tp {
	default:
		return new(app.HTTPServer)
	}
}
