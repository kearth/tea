package out

import (
	"example/local/app/infra/ecode"

	"github.com/kearth/klib/kerr"
	"github.com/kearth/tea/frame/base"
	"github.com/kearth/tea/frame/server/httpserver"
)

// WriteString 写入字符串
func WriteString(r *httpserver.Request, data string) {
	r.Response.WriteString(data)
}

// Flush 刷新响应
func Flush(r *httpserver.Request) {
	r.Response.Flush()
}

// JSONResponse JSON响应
func JSONResponse(r *httpserver.Request, err kerr.Error, data any) {
	if err == nil {
		err = ecode.Succ
	}
	out := base.Response()
	out.SetCode(err.Code())
	out.SetMsg(err.Display())
	out.SetData(data)
	r.Response.WriteJson(out.String())
	r.Exit()
}

// FileResponse 文件响应
func FileResponse(r *httpserver.Request, filePath string) {
	r.Response.ServeFileDownload(filePath)
	r.Exit()
}

// HTMLResponse HTML响应
func HTMLResponse(r *httpserver.Request, html string) {
	r.Header.Set("Content-Type", "text/html; charset=utf-8")
	r.Response.WriteString(html)
	r.Exit()
}
