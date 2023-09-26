package tea

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/kearth/tea/core"
	"github.com/spf13/cast"
)

// Check Interface
var _ core.IContainer = &HTTPServer{}

// Bootstrap
type Bootstrap func(ctx context.Context) error
type RouterFunc func(ctx context.Context) *HTTPRouter

// init
func init() {
	// register
	core.IOC().Register(new(HTTPServer))
}

// HTTPServer
type HTTPServer struct {
	http.Server
	ConfigPath    string
	BootstrapFunc Bootstrap
	RouterFunc    RouterFunc
}

// Name
func (h *HTTPServer) Name() string {
	return "HTTPServer"
}

// New
func (h *HTTPServer) New() core.IContainer {
	// Default
	h.ConfigPath = "./conf/app.toml"
	h.RouterFunc = func(ctx context.Context) *HTTPRouter {
		defaultRouter := NewHTTPRouter()
		defaultRouter.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, "hello, world!\n")
		})
		return defaultRouter
	}
	h.BootstrapFunc = func(ctx context.Context) error {
		return nil
	}
	return h
}

// Init
func (h *HTTPServer) Init(ctx context.Context) error {
	var err error
	httpconfig, err := core.IOC().Get("HTTPConfig")
	config := httpconfig.(*HTTPConfig)
	if err = core.Parse(h.ConfigPath, &config); err != nil {
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
		return err
	}
	h.Handler = h.RouterFunc(ctx)
	//TODO
	fmt.Println("ListenAndServe:8080")
	return h.Server.ListenAndServe()
}

// Shutdown
func (h *HTTPServer) Shutdown(ctx context.Context, cancel context.CancelFunc) error {
	// TODO
	return nil
}

// Bootstrap
func (h *HTTPServer) SetBootstrap(bs Bootstrap) {
	h.BootstrapFunc = bs
}

// SetRouter
func (h *HTTPServer) SetRouter(r RouterFunc) {
	h.RouterFunc = r
}

// SetConf
func (h *HTTPServer) SetConf(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	h.ConfigPath = path
	return nil
}
