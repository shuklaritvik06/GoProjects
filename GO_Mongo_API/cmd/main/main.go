package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/GO_Mongo_API/pkg/routes"
)

func main() {
	router := &mux.Router{}
	routes.EmployeeRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
