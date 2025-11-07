package api

import (
	"github.com/kearth/tea/frame/server/httpserver"
)

// Hello
type Hello struct {
}

// Say 打招呼
func (c *Hello) Say(r *httpserver.Request) {
	r.Response.Write("hello world")
}
