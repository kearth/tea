package router

import (
	"example/local/controllers"

	"github.com/kearth/tea"
)

func Router(ctx tea.Context) *tea.HTTPRouter {
	router := tea.NewHTTPRouter()
	router.Get("/api/who", controllers.Hello)
	return router
}
