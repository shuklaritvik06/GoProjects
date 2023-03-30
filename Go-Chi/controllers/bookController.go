package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shuklaritvik06/GoProjects/gochi/database"
	"github.com/shuklaritvik06/GoProjects/gochi/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	cursor, _ := database.GetDB().Database("book").Collection("books").Find(context.Background(), bson.D{})
	var books []models.Book
	if err := cursor.All(context.TODO(), &books); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	bookName := chi.URLParam(r, "bookName")
	filter := bson.D{{Key: "title", Value: bookName}}
	err := database.GetDB().Database("book").Collection("books").FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	_, err := database.GetDB().Database("book").Collection("books").InsertOne(context.Background(), book)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	filter := bson.D{{Key: "title", Value: book.Title}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "author", Value: book.Author},
			{Key: "publication", Value: book.Publication},
			{Key: "title", Value: book.Title},
		}},
	}
	_, err := database.GetDB().Database("book").Collection("books").UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	_, err := database.GetDB().Database("book").Collection("books").DeleteMany(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("All books deleted")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookName := chi.URLParam(r, "bookName")
	filter := bson.D{{Key: "title", Value: bookName}}
	_, err := database.GetDB().Database("book").Collection("books").DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Book deleted")
}
