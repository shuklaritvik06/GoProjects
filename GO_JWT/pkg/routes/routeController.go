package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/GO_JWT/pkg/controllers"
)

var JWTRoutes = func(router *mux.Router) {
	router.HandleFunc("/register/jwt", controllers.JWTRegister).Methods("POST")
	router.HandleFunc("/login/jwt", controllers.JWTLogin).Methods("POST")
}
