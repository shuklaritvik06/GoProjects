package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/GO_AUTH/pkg/controllers"
)

var OauthRoutes = func(router *mux.Router) {
	router.HandleFunc("/login/google", controllers.GoogleLogin).Methods("GET")
	router.HandleFunc("/login/google/callback", controllers.GoogleCallback).Methods("GET")
}
