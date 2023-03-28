package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/filecompress/backend/config"
	"github.com/shuklaritvik06/GoProjects/filecompress/backend/routes"
)

func main() {
	config.Configure()
	router := mux.NewRouter()
	router.PathPrefix("/api/v1")
	routes.AuthRoutes(router)
	routes.FileRoutes(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
