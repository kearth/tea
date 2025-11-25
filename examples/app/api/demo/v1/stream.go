package v1

import (
	"bufio"
	"example/local/app/infra/ecode"
	"example/local/app/infra/out"
	"example/local/app/infra/structs"
	"fmt"

	"github.com/kearth/klib/kerr"
	"github.com/kearth/tea/frame/base"

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
		if len(buffer) == 0 {
			fmt.Println("empty line")
			break
		}
		demo := &structs.Demo{
			Type: "text",
			Text: string(buffer),
		}
		out.WriteString(r, "data: "+demo.String()+"\n\n")
		out.Flush(r)
	}
	// 检查扫描错误
	o := base.Response()
	var e kerr.Error
	if err := scanner.Err(); err != nil {
		e = ecode.ScannerError.Wrap(err)
	} else {
		e = ecode.Succ
	}
	o.SetCode(e.Code())
	o.SetMsg(e.Display())
	o.SetData(nil)
	out.WriteString(r, o.String())
	out.Flush(r)
}
