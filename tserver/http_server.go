package tserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kearth/tea/app"
	"github.com/kearth/tea/conf"
	"github.com/spf13/cast"
)

// IHTTPRouter
type IHTTPRouter interface {
	http.Handler
	app.IRouter
}

type HTTPRouter struct {
	http.ServeMux
}

func (h *HTTPRouter) Group(pattern string) app.IRouter { return h }
func (h *HTTPRouter) Use()                             {}

//type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// HTTPServer
type HTTPServer struct {
	http.Server
	ConfigPath    string
	BootstrapFunc app.Bootstrap
	Router        IHTTPRouter
}

// HTTPConfig
type HTTPConfig struct {
	Name         string `toml:"Name"`
	Port         int    `toml:"Port"`
	ReadTimeout  int    `toml:"ReadTimeout"`
	WriteTimeout int    `toml:"WriteTimeout"`
	IdleTimeout  int    `toml:"IdleTimeout"`
}

// Init
func (h *HTTPServer) Init(ctx context.Context) error {
	var (
		config = new(HTTPConfig)
		err    error
	)
	if err = conf.Parse(h.ConfigPath, &config); err != nil {
		panic(err)
	}
	if config.Port != 0 {
		h.Addr = fmt.Sprintf(":%s", cast.ToString(config.Port))
	}
	var to int
	if to = config.ReadTimeout; to > 0 {
		h.ReadTimeout = time.Duration(to) * time.Millisecond
	}
	if to = config.WriteTimeout; to > 0 {
		h.WriteTimeout = time.Duration(to) * time.Millisecond
	}
	if to = config.IdleTimeout; to > 0 {
		h.IdleTimeout = time.Duration(to) * time.Millisecond
	}
	return nil
}

// Start
func (h *HTTPServer) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer h.Shutdown(ctx, cancel)
	if err := h.Init(ctx); err != nil {
		return err
	}
	if err := h.BootstrapFunc(ctx); err != nil {
	}
	// TODO
	h.Handler = h.Router
	fmt.Println("ListenAndServe:8080")
	return h.Server.ListenAndServe()
}

// Shutdown
func (h *HTTPServer) Shutdown(ctx context.Context, cancel context.CancelFunc) error {
	return nil
}

// Bootstrap
func (h *HTTPServer) SetBootstrap(bs app.Bootstrap) error {
	h.BootstrapFunc = bs
	// TODO
	return nil
}

// TODO
// RegisetrRouter
func (h *HTTPServer) SetRouter(r app.IRouter) error {
	h.Router = r.(IHTTPRouter)
	// TODO
	return nil
}

// RegisetrRouter
func (h *HTTPServer) SetConf(path string) error {
	h.ConfigPath = path
	// TODO
	return nil
}

func NewRouter() *HTTPRouter {
	return new(HTTPRouter)
}

// DefaultServer
/*
func (h *HTTPServer) DefaultServer() *http.Server {
	h.Handler = (http.HandlerFunc)(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	return &h.Server
}
*/
