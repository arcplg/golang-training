package db

import (
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func Connect() {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	port := os.Getenv("MONGO_URI")
	host := os.Getenv("MONGO_HOST")

	client, err := mongo.Connect(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	MongoClient = client

}

func GetCollection(collection string) *mongo.Collection {
	database := os.Getenv("MONGO_DATABASE")
	return MongoClient.Database(database).Collection(collection)
}
