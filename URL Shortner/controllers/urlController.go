package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/shuklaritvik06/GoProjects/urlshortner/database"
	"github.com/shuklaritvik06/GoProjects/urlshortner/helpers"
	"github.com/shuklaritvik06/GoProjects/urlshortner/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var url models.URL
	_ = json.NewDecoder(r.Body).Decode(&url)
	url.Created_Date = time.Now()
	newurl := helpers.ShortenURL(url.URL, r)
	newdoc := models.URL{
		URL:          url.URL,
		Short_URL:    newurl,
		Created_Date: time.Now(),
	}
	database.GetDB().Database("url").Collection("url").InsertOne(context.Background(), newdoc)
	w.Write([]byte(newurl))
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	var url models.URL
	var urldoc models.URL
	_ = json.NewDecoder(r.Body).Decode(&url)
	result := database.GetDB().Database("url").Collection("url").FindOne(context.Background(), bson.D{{Key: "short_url", Value: url.Short_URL}})
	result.Decode(&urldoc)
	w.Write([]byte(urldoc.URL))
}
