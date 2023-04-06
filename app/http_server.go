package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kearth/tea/conf"
	"github.com/spf13/cast"
)

// HTTPServer
type HTTPServer struct {
	http.Server
}

// HTTPConfig
type HTTPConfig struct {
	AppName    string
	RunMode    string
	Port       int
	IDC        string
	HttpServer *struct {
		ReadTimeout  int
		WriteTimeout int
		IdleTimeout  int
	}
}

func (c *HTTPConfig) Parse() error {
	if err := conf.Parse(path, &appConf); err != nil {
		return nil, err
	}
	return appConf, nil
}

// Init
func (h *HTTPServer) Init(ctx context.Context, appConfig *AppConfig) {
	config := appConfig.(HTTPConfig)
	if config.Port != 0 {
		h.Addr = fmt.Sprintf(":%s", cast.ToString(config.Port))
	}
	if config.HttpServer != nil {
		var to int
		if to = config.HttpServer.ReadTimeout; to > 0 {
			h.ReadTimeout = time.Duration(to) * time.Millisecond
		}
		if to = config.HttpServer.WriteTimeout; to > 0 {
			h.WriteTimeout = time.Duration(to) * time.Millisecond
		}
		if to = config.HttpServer.IdleTimeout; to > 0 {
			h.IdleTimeout = time.Duration(to) * time.Millisecond
		}
	}
}

// LoadConfig
func (h *HTTPServer) LoadConfig(configPath string) *AppConfig {
	var (
		appConfig *HTTPConfig
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
