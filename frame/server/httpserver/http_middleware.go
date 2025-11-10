package httpserver

import "github.com/gogf/gf/v2/net/ghttp"

// Middleware HTTP 中间件
type Middleware = func(r *Request)

// MiddlewareHandlerResponse HTTP 中间件 - 响应处理
func MiddlewareHandlerResponse() Middleware {
	return ghttp.MiddlewareHandlerResponse
}
