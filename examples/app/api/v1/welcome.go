package v1

/**
 * Welcome 欢迎页
 */

import (
	"fmt"

	"github.com/kearth/tea/frame/server"
	"github.com/kearth/tea/frame/server/httpserver"
)

type Welcome struct {
}

func (w *Welcome) Welcome(r *httpserver.Request) {
	r.Header.Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(r.Response.Writer, server.GetWelcome())
}
