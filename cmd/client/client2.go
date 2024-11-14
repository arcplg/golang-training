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

	chatWithUser1(client)
}

func chatWithUser1(c handlers.ChatAppServiceClient) {
	fmt.Println("Chat with user...!")

	//waitc := make(chan struct{})

	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE2NDI4MDgsIm5iZiI6MTczMTU1NjQwOCwiaWF0IjoxNzMxNTU2NDA4LCJqdGkiOiJlMDU1YzI5Ny01MThhLTQzMzAtYjA2Zi1hMzZmOWJhZGUxZWEiLCJ1c2VybmFtZSI6Imt1bmgwemRuIiwiaWQiOiJlMDU1YzI5Ny01MThhLTQzMzAtYjA2Zi1hMzZmOWJhZGUxZWEifQ.EPVSwbbVO1Hbj9QvmYLnVmowVZS3mDA1214CtYrTzH0"
	md := metadata.Pairs(
		"authorization", "Bearer "+bearerToken,
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, errChatClient := c.ChatWithUserStream(ctx)

	if errChatClient != nil {
		log.Fatalf("error connection %v", errChatClient)
	}

	go func() {
		//var mess string
		//log.Println("Enter Message")
		//
		//fmt.Scan(&mess)

		messages := []handlers.ChatWithUserMessage{
			handlers.ChatWithUserMessage{
				Message:  "mess",
				Receiver: "24eadac5-6216-4131-91bc-4e53cd426bcf",
				Sender:   "e055c297-518a-4330-b06f-a36f9bade1ea",
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
