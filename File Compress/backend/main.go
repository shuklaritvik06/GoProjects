package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/shuklaritvik06/GoProjects/filecompress/backend/config"
	"github.com/shuklaritvik06/GoProjects/filecompress/backend/routes"
)

func main() {
	config.Configure()
	router := mux.NewRouter()
	routes.AuthRoutes(router)
	routes.FileRoutes(router)
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
