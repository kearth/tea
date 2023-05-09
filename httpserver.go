package tea

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kearth/tea/conf"
	"github.com/spf13/cast"
)

// IHTTPRouter
type IHTTPRouter interface {
	http.Handler
	IRouter
}

// HTTPRouter
type HTTPRouter struct {
	http.ServeMux
}

func (h *HTTPRouter) Group(pattern string) IRouter { return h }
func (h *HTTPRouter) Use()                         {}

//type HandlerFunc func(w http.ResponseWriter, r *http.Request)

var _ IContainer = &HTTPServer{}

// init
func init() {
	// register
	IOC().Register(new(HTTPServer))
}

// HTTPServer
type HTTPServer struct {
	http.Server
	ConfigPath    string
	BootstrapFunc Bootstrap
	Router        IHTTPRouter
}

// Name
func (h *HTTPServer) Name() string {
	return "HTTPServer"
}

// New
func (h *HTTPServer) New() IContainer {
	// Default
	h.ConfigPath = "./conf/app.toml"
	return h
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
func (h *HTTPServer) SetBootstrap(bs Bootstrap) {
	h.BootstrapFunc = bs
}

// TODO
// RegisetrRouter
func (h *HTTPServer) SetRouter(r IRouter) error {
	h.Router = r.(IHTTPRouter)
	// TODO
	return nil
}

// SetConf
func (h *HTTPServer) SetConf(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	h.ConfigPath = path
	return nil
}

// NewRouter
func NewRouter() *HTTPRouter {
	return new(HTTPRouter)
}
