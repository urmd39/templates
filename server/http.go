package server

import (
	"net/http"
	"strings"
	"templates/infrastructure"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router() http.Handler {
	r := chi.NewRouter()

	basePath := infrastructure.BasePath
	r.Get(basePath+"/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(infrastructure.HttpSwagger), //The url pointing to API definition"
	))

	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Homepage"))
	//})

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
