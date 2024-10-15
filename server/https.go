package server

import (
	"net/http"
	"strings"
	"templates/infrastructure"
	"templates/middleware"

	httpSwagger "github.com/swaggo/http-swagger"

	// _ "backend/docs"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(chiMiddleware.DefaultLogger)

	myLogger := infrastructure.InitLoggerWithNATSHook()
	r.Use(middleware.NewStructuredLogger(myLogger))

	basePath := infrastructure.BasePath
	r.Get(basePath+"/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(infrastructure.HttpSwagger), //The url pointing to API definition"
	))

	r.Route(basePath, func(subR chi.Router) {
	})
	return r
}

// HandleHTTPServer
func HandleHTTPServer(port string) {
	infrastructure.InfoLog.Printf("server  %s", port)
	infrastructure.InfoLog.Printf("swagger : %s", strings.Replace(infrastructure.HttpSwagger, "doc.json", "index.html", 1))

	if err := http.ListenAndServe(port, Router()); nil != err {
		infrastructure.ErrLog.Print(err)
	}
}
