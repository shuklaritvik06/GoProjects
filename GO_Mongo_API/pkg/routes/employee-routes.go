package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/GO_Mongo_API/pkg/controllers"
)

var EmployeeRoutes = func(router *mux.Router) {
	router.HandleFunc("/create/", controllers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/", controllers.GetEmployee).Methods("GET")
	router.HandleFunc("/employee/{id}", controllers.GetEmployeeByID).Methods("GET")
	router.HandleFunc("/employee/{id}", controllers.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", controllers.DeleteEmployee).Methods("DELETE")
}
