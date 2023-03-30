package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/shuklaritvik06/GoProjects/gochi/controllers"
)

func BookRoutes(r *chi.Mux) {
	r.Get("/getbooks", controllers.GetBooks)
	r.Get("/book/{bookName}", controllers.GetBook)
	r.Post("/addbook", controllers.AddBook)
	r.Post("/updatebook/{bookName}", controllers.UpdateBook)
	r.Delete("/deletebooks", controllers.DeleteBooks)
	r.Delete("/delete/{bookName}", controllers.DeleteBook)
}
