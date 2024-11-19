package main

import (
	"context"
	"log"
	"tracking-learning/config"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	opt := option.WithCredentialsFile(config.Firebase.ServiceAccountKeyPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v\n", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing Firebase Auth: %v\n", err)
	}

	log.Println("Firebase connected successfully")
	_ = authClient
}
