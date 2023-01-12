package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c *mongo.Client

func Configure() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<USERNAME>:<PASSWORD>@cluster0.aeiaykn.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	c = client
}

func GetClient() *mongo.Client {
	return c
}
