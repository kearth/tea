package v1

/**
 * Hello 打招呼
 */

import (
	"context"
)

// Hello
type Hello struct{}

// HelloReq 请求参数
type HelloReq struct {
	Name string `json:"name"`
}

// HelloRes 响应参数
type HelloRes struct {
	Message string `json:"message"`
}

// Say 打招呼
func (c *Hello) Hello(ctx context.Context, h *HelloReq) (*HelloRes, error) {
	return &HelloRes{
		Message: "hello " + h.Name,
	}, nil
}
