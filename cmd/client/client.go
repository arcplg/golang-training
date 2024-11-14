package main

import (
	"github.com/ngocthanh06/chatapp/internal/handlers"
	"github.com/ngocthanh06/chatapp/internal/storage"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50070", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("ip exist", err)
	}

	defer conn.Close()

	client := handlers.NewChatAppServiceClient(conn)
	log.Println("Connection client success...")

	storage.New()
	ChatWithUser(client, "24eadac5-6216-4131-91bc-4e53cd426bcf")
}
