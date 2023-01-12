package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/GO_MySQL_API/pkg/models"
	"github.com/shuklaritvik06/GoProjects/GO_MySQL_API/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	id, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		log.Fatal("Error in converting the string id")
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b, err := book.CreateBook()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	ID := vars["id"]
	id, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		log.Fatal("Error in converting the string id")
	}
	book, db := models.GetBookById(id)
	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}
	if updatedBook.Name != "" {
		book.Name = updatedBook.Name
	}
	if updatedBook.Publication != "" {
		book.Publication = updatedBook.Publication
	}
	if updatedBook.Price != 0 {
		book.Price = updatedBook.Price
	}
	db.Save(book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	id, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		log.Fatal("Error in converting the string id")
	}
	models.DeleteBook(id)
	res, _ := json.Marshal("Record Successfully Deleted!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
