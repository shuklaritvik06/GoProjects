package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shuklaritvik06/GoProjects/gochi/config"
	"github.com/shuklaritvik06/GoProjects/gochi/routes"
)

func main() {
	config.Configure()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	routes.BookRoutes(r)
	http.ListenAndServe(":3000", r)
}
