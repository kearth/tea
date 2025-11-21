package v1

import (
	"example/local/app/infra/out"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server/httpserver"
)

// Stream 页面 - 流响应
type Stream struct{}

func (w *Stream) Chat(r *httpserver.Request) {
	html := gfile.GetContents(env.GetResourcesDir() + "/public/html/demo/stream.html")
	out.HTMLResponse(r, html)
}
