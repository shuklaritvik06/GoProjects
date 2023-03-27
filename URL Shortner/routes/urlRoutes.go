package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/urlshortner/controllers"
)

func URLRoutes(r *mux.Router) {
	r.HandleFunc("/shorten", controllers.ShortenURL).Methods("POST")
	r.HandleFunc("/real", controllers.RedirectURL).Methods("GET")
}
