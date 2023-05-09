package tea

import (
	"net/http"
)

var _ IContainer = &HTTPRouter{}

// HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request)

// MiddleWareFunc
type MiddleWareFunc struct{}

// HTTPRouter
type HTTPRouter struct {
	http.ServeMux
}

// init
func init() {
	// register
	IOC().Register(new(HTTPRouter))
}

// Name
func (h *HTTPRouter) Name() string {
	return "HTTPRouter"
}

// New
func (h *HTTPRouter) New() IContainer {
	return h
}

// NewHTTPRouter
func NewHTTPRouter() *HTTPRouter {
	httpRouter, err := IOC().Get("HTTPRouter")
	if err != nil {
		panic(err)
	}
	return httpRouter.(*HTTPRouter)
}

/*
func (h *HTTPRouter) Group(pattern string) IRouter {
	return h
}
*/

// Get
func (h *HTTPRouter) Get(pattern string, hfc HandlerFunc) {
	h.HandleFunc(pattern, func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			hfc(rw, r)
		}
	})
}

// Post
func (h *HTTPRouter) Post(pattern string, hfc HandlerFunc) {
	h.HandleFunc(pattern, func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			hfc(rw, r)
		}
	})
}

// Any
func (h *HTTPRouter) Any(pattern string, hfc HandlerFunc) {
	h.HandleFunc(pattern, func(rw http.ResponseWriter, r *http.Request) {
		hfc(rw, r)
	})
}
