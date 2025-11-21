package v1

import (
	"bufio"
	"example/local/app/infra/ecode"
	"example/local/app/infra/out"
	"example/local/app/infra/structs"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/kearth/tea/frame/env"
	"github.com/kearth/tea/frame/server/httpserver"
)

/**
 * Stream 流式响应
 */
type Stream struct{}

// Chat 流式响应
func (s *Stream) Chat(r *httpserver.Request) {

	// 流式响应头设置
	out.StreamHeader(r)

	// 打开文件
	file, err := gfile.Open(env.GetResourcesDir() + "/public/file/demo/stream.md")
	if err != nil {
		out.JSONResponse(r, ecode.FileNotFound.Wrap(err), nil)
		return
	}
	defer file.Close()

	// 读取文件内容
	scanner := bufio.NewScanner(file)
	var buffer []byte
	for scanner.Scan() {
		buffer = scanner.Bytes()
		demo := &structs.Demo{
			Type: "text",
			Text: string(buffer),
		}
		out.WriteString(r, "data: "+demo.String()+"\n\n")
		out.Flush(r)
	}
	// 检查扫描错误
	if err := scanner.Err(); err != nil {
		out.JSONResponse(r, ecode.ScannerError.Wrap(err), nil)
		return
	}
	// 响应成功
	out.JSONResponse(r, ecode.Succ, nil)
}
