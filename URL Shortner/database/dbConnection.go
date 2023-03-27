package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func Database() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DB_URI")).SetServerAPIOptions(serverAPI)
	client, _ := mongo.Connect(context.TODO(), opts)
	db = client
	fmt.Println("Successfully connected to MongoDB!")
}

func GetDB() *mongo.Client {
	return db
}
