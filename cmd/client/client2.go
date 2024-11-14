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
	ChatWithUser(client, "e055c297-518a-4330-b06f-a36f9bade1ea")
}
