package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/shuklaritvik06/GoProjects/GO_AUTH/pkg/config"
	"github.com/shuklaritvik06/GoProjects/GO_AUTH/pkg/routes"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	config.Configure()
	router := &mux.Router{}
	routes.OauthRoutes(router)
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
