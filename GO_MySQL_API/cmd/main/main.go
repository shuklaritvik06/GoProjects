package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/GO_MySQL_API/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.BookStoreRoutes(router)
	log.Fatal(http.ListenAndServe("127.0.0.1:5000", router))
}
