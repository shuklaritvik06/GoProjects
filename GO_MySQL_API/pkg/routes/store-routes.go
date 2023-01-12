package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/GO_MySQL_API/pkg/controllers"
)

var BookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}/", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}/", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}/", controllers.DeleteBook).Methods("DELETE")
}
