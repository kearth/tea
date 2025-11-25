package out

import "github.com/gogf/gf/v2/net/ghttp"

// StreamHeader 流式响应头设置
func StreamHeader(r *ghttp.Request) {
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("Transfer-Encoding", "chunked")
}
