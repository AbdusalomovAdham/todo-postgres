package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDb() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal("Error connect mongodb", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Mongodb ping error", err)
	}
	DB = client.Database("todo-app")
	log.Println("Mongo connected seccessfully!")
	return DB
}
