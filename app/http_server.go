package app

import (
	"context"
	"time"
)

type HTTPServer struct {
	Ctx          context.Context
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (h *HTTPServer) Init()        {}
func (h *HTTPServer) LoadConfig()  {}
func (h *HTTPServer) Start() error {}
