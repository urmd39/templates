package api

import (
	"templates/controller"
	"templates/middleware"

	chi "github.com/go-chi/chi/v5"
)

func NewApiRouter(subR chi.Router) {
	authorizer := middleware.NewAuthentication()
	demoController := controller.NewDemoController()

	// Private routes
	subR.Group(func(r chi.Router) {
		r.Use(authorizer.Authorizer())
	})

	// Public routes
	subR.Group(func(r chi.Router) {

	})
}
