package main

import (
	"context"
	"fmt"
	"github.com/ngocthanh06/chatapp/internal/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
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

	chatWithUser(client)
}

func chatWithUser(c handlers.ChatAppServiceClient) {
	fmt.Println("Chat with user...!")

	//waitc := make(chan struct{})

	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE2Mzk3MDgsIm5iZiI6MTczMTU1MzMwOCwiaWF0IjoxNzMxNTUzMzA4LCJqdGkiOiIyNGVhZGFjNS02MjE2LTQxMzEtOTFiYy00ZTUzY2Q0MjZiY2YiLCJ1c2VybmFtZSI6InRoYW5oZG4iLCJpZCI6IjI0ZWFkYWM1LTYyMTYtNDEzMS05MWJjLTRlNTNjZDQyNmJjZiJ9.v9FweAzH5ohpqQDqM5_gzHO8pyceQ8a9W1geu6i8EmI"
	md := metadata.Pairs(
		"authorization", "Bearer "+bearerToken,
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, errChatClient := c.ChatWithUserStream(ctx)

	if errChatClient != nil {
		log.Fatalf("error connection %v", errChatClient)
	}

	go func() {
		messages := []handlers.ChatWithUserMessage{
			handlers.ChatWithUserMessage{
				Message:  "Hello fen thanh client",
				Receiver: "e055c297-518a-4330-b06f-a36f9bade1ea",
				Sender:   "24eadac5-6216-4131-91bc-4e53cd426bcf",
			},
			//handlers.ChatWithUserMessage{
			//	Message:  "Hello fen thanh dn",
			//	Receiver: "e055c297-518a-4330-b06f-a36f9bade1ea",
			//	Sender:   "24eadac5-6216-4131-91bc-4e53cd426bcf",
			//},
		}

		for _, req := range messages {
			if errReq := stream.Send(&req); errReq != nil {
				log.Fatalf("Error send message %v", errReq)
			}
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			resp, errRcv := stream.Recv()
			if errRcv == io.EOF {
				log.Println("End of file")
				break
			}

			if errRcv != nil {
				break
			}

			log.Printf("resp %v", resp.GetMessage())
		}

		//close(waitc)
	}()

	select {}
	//<-waitc
}
