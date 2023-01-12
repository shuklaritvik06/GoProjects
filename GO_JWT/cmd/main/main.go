package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/GO_JWT/pkg/config"
	"github.com/shuklaritvik06/GoProjects/GO_JWT/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.JWTRoutes(router)
	config.Configure()
	log.Fatal(http.ListenAndServe(":5000", router))
}
