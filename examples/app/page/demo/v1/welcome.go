package v1

/**
 * Welcome 欢迎页
 */

import (
	"example/local/app/infra/out"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server/httpserver"
)

// Welcome 欢迎页
type Welcome struct{}

func (w *Welcome) Welcome(r *httpserver.Request) {
	html := gfile.GetContents(env.GetResourcesDir() + "/public/html/demo/welcome_v1.html")
	out.HTMLResponse(r, html)
}
