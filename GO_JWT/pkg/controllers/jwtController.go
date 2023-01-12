package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/shuklaritvik06/GoProjects/GO_JWT/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

type Response struct {
	Token string
}

type Input struct {
	Username string
	Password string
}

type User struct {
	Username string `bson:"username,omitempty"`
	Token    string `bson:"token,omitempty"`
	Password string `bson:"password,omitempty"`
}

func JWTRegister(w http.ResponseWriter, r *http.Request) {
	db = config.GetClient()
	username := r.FormValue("username")
	password := r.FormValue("password")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"nbf":      time.Date(2022, 12, 13, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	data := User{
		Username: username,
		Token:    tokenString,
		Password: password,
	}
	db.Database("user").Collection("users").InsertOne(context.Background(), data)
	resp, _ := json.Marshal(Response{
		Token: tokenString,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func JWTLogin(w http.ResponseWriter, r *http.Request) {
	db = config.GetClient()
	params := r.URL.Query()
	res := db.Database("user").Collection("users").FindOne(context.TODO(), bson.D{
		{Key: "username", Value: params["username"]},
	})
	if res == nil {
		log.Fatal("No user found")
	}
	token, err := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["username"] != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Login Successful"))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Login Failed"))
		}
	} else {
		fmt.Println(err)
	}
}
