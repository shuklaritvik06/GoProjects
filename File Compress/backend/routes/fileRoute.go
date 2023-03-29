package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/filecompress/backend/controllers"
)

func FileRoutes(r *mux.Router) {
	r.HandleFunc("/compress", controllers.Compress).Methods("POST", "OPTIONS")
	r.HandleFunc("/decompress", controllers.Decompress).Methods("POST", "OPTIONS")
}
