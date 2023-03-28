package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/filecompress/backend/controllers"
	"github.com/shuklaritvik06/GoProjects/filecompress/backend/middleware"
)

func FileRoutes(r *mux.Router) {
	r.HandleFunc("/compress", controllers.Compress).Methods("POST")
	http.Handle("/compress", middleware.AuthMiddleware(r))
	r.HandleFunc("/decompress", controllers.Decompress).Methods("POST")
	http.Handle("/decompress", middleware.AuthMiddleware(r))
}
