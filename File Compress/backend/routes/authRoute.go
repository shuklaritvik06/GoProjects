package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/filecompress/backend/controllers"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.Signup).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/refresh", controllers.Refresh).Methods("POST", "OPTIONS")
}
