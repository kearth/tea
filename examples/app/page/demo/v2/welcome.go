package v2

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server/httpserver"
)

type Welcome struct{}

func (w *Welcome) Welcome(r *httpserver.Request) {
	r.Header.Set("Content-Type", "text/html; charset=utf-8")
	html := gfile.GetContents(env.GetResourcesDir() + "/public/html/demo/welcome_v2.html")
	r.Response.WriteString(html)
}
