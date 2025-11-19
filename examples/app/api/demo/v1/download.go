package v1

import (
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server/httpserver"
)

/*
 * File 下载文件
 */

type File struct{}

// Download 下载文件
func (c *File) Image(r *httpserver.Request) {
	r.Response.ServeFileDownload(env.GetResourcesDir() + "/public/image/demo/tea.png")
}
