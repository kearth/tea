package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cast"
)

// HTTPServer
type HTTPServer struct {
	http.Server
}

// HTTPSServer
// TODO
type HTTPSServer struct {
	http.Server
}

// Init
func (h *HTTPServer) Init(ctx context.Context, appConfig *AppConfig) {
	if appConfig.Port != 0 {
		h.Addr = fmt.Sprintf(":%s", cast.ToString(appConfig.Port))
	}
	if appConfig.HttpServer != nil {
		var to int
		if to = appConfig.HttpServer.ReadTimeout; to > 0 {
			h.ReadTimeout = time.Duration(to) * time.Millisecond
		}
		if to = appConfig.HttpServer.WriteTimeout; to > 0 {
			h.WriteTimeout = time.Duration(to) * time.Millisecond
		}
		if to = appConfig.HttpServer.IdleTimeout; to > 0 {
			h.IdleTimeout = time.Duration(to) * time.Millisecond
		}
	}
}

// LoadConfig
func (h *HTTPServer) LoadConfig(configPath string) *AppConfig {
	var (
		appConfig *AppConfig
		err       error
	)
	if appConfig, err = LoadAppConfig(configPath); err != nil {
		panic(err)
	}
	return appConfig
}

// Start
func (h *HTTPServer) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	appConfig := h.LoadConfig(GetAppConfigPath())
	h.Init(ctx, appConfig)
	return h.Listen()
}

// Listen
func (h *HTTPServer) Listen() error {
	s := h.DefaultServer()
	fmt.Println("ListenAndServe:8080")
	return s.ListenAndServe()
}

// DefaultServer
func (h *HTTPServer) DefaultServer() *http.Server {
	h.Handler = (http.HandlerFunc)(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	return &h.Server
}
