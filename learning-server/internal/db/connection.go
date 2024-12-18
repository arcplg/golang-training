package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func ConnectDB() *mongo.Database {

	mongoUrl := os.Getenv("MONGO_URL")
	database := os.Getenv("MONGO_DATABASE")

	fmt.Println(mongoUrl)

	client, err := mongo.Connect(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client.Database(database)
}

func GetCollection(collection string) *mongo.Collection {
	return ConnectDB().Collection(collection)
}
