package httpserver

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// 自定义 MIME 映射（补充 GF 内置未覆盖的后缀）
var customMimeMap = map[string]string{
	".md":   "text/markdown; charset=utf-8",
	".yml":  "application/yaml; charset=utf-8",
	".toml": "application/toml; charset=utf-8",
	".json": "application/json; charset=utf-8",
}

// bindStaticFile 绑定静态文件
func bindStaticFile(root string, sf map[string]string) map[string]any {
	sfmap := make(map[string]any)
	for uri, file := range sf {
		path := gfile.Join(root, file)
		if !gfile.IsFile(path) {
			continue
		}
		key := fmt.Sprintf("%s:%s", GET, uri)
		sfmap[key] = func(r *Request) {
			file := gfile.GetContents(path)
			// 1. 提取文件后缀（含 "."，如 ".html"）
			ext := gfile.Ext(path)
			// 统一后缀为小写（避免 ".HTML" 和 ".html" 识别不一致）
			ext = gstr.ToLower(ext)

			// 2. 确定 MIME 类型（优先自定义映射，再用 GF 内置）
			mimeType := ""
			if customType, ok := customMimeMap[ext]; ok {
				mimeType = customType
			} else {
				// GF 内置 MIME 映射（覆盖大部分常见文件：html、css、js、png、jpg 等）
				// mimeType = ghttp.MIME(ext)
				// 兜底：若 GF 也无法识别，默认设为二进制流（避免浏览器下载错误）
				if mimeType == "" {
					mimeType = "application/octet-stream"
				}
				// 文本类文件统一添加 utf-8 编码（避免中文乱码）
				if gstr.InArray([]string{
					"text/plain", "text/css", "text/javascript",
					"text/html", "application/xml",
				}, mimeType) {
					mimeType += "; charset=utf-8"
				}
			}

			// 4. 设置响应头并返回内容
			r.Header.Set("Content-Type", mimeType)
			r.Response.WriteString(file)
			r.Exit()
		}
	}
	return sfmap
}
