package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/urlshortner/config"
	"github.com/shuklaritvik06/GoProjects/urlshortner/database"
	"github.com/shuklaritvik06/GoProjects/urlshortner/routes"
)

func main() {
	config.Configure()
	database.Database()
	router := mux.NewRouter()
	routes.URLRoutes(router)
	http.ListenAndServe(":8000", router)
}
