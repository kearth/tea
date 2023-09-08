package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/kearth/tea"
)

func main() {
	httpserver := tea.NewHTTPServer()
	httpserver.SetBootstrap(Bootstrap)
	httpserver.SetRouter(Router)
	httpserver.Start()
}

func Bootstrap(ctx context.Context) error {
	fmt.Println("hello world")
	return nil
}

func Router(ctx context.Context) *tea.HTTPRouter {
	router := tea.NewHTTPRouter()
	router.Get("/api/who", Who)
	return router
}

func Who(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "666\n")
}
