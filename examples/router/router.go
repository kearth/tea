package router

import (
	"context"
	"example/local/controllers"

	"github.com/kearth/tea"
)

func Router(ctx context.Context) *tea.HTTPRouter {
	router := tea.NewHTTPRouter()
	router.Get("/api/who", controllers.Hello)
	return router
}
