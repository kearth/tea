package v1

import (
	"example/local/app/infra/out"

	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server/httpserver"
)

/*
 * File 下载文件
 */

type File struct{}

// Image 下载图片
func (c *File) Image(r *httpserver.Request) {
	out.FileResponse(r, env.GetResourcesDir()+"/public/image/demo/tea.png")
}
