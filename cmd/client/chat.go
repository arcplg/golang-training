package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ngocthanh06/chatapp/internal/handlers"
	"github.com/ngocthanh06/chatapp/internal/middleware"
	"github.com/ngocthanh06/chatapp/internal/models"
	"github.com/ngocthanh06/chatapp/internal/storage"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	"time"
)

func ChatWithUser(c handlers.ChatAppServiceClient, receiver string) {
	fmt.Println("Chat with user...!")

	var username string
	var password string
	var user models.User
	var bearerToken string
	count := 1

	for count < 5 {
		log.Println("Please enter username:")
		fmt.Scan(&username)

		log.Println("Please enter password:")
		fmt.Scan(&password)

		// login function
		result := storage.GetDB().Table("users").Where("username", username).First(&user)

		if result.RowsAffected == 0 {
			log.Println("Username or password is correct")
			count++
			continue
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

			if err != nil {
				log.Printf("Username or password is wrong: %v\n", err)

				continue
			}

			claims := middleware.Claims{
				jwt.RegisteredClaims{
					// A usual scenario is to set the expiration time relative to the current time
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
					IssuedAt:  jwt.NewNumericDate(time.Now()),
					NotBefore: jwt.NewNumericDate(time.Now()),
					ID:        user.Id.String(),
				},
				user.Username,
				user.Id.String(),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			bearerToken, _ = token.SignedString([]byte(os.Getenv("JWT_KEY")))

			break
		}
	}

	if count == 5 {
		log.Println("Please re-login after")

		return
	}

	md := metadata.Pairs(
		"authorization", "Bearer "+bearerToken,
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, errChatClient := c.ChatWithUserStream(ctx)

	if errChatClient != nil {
		log.Fatalf("error connection %v", errChatClient)
	}

	go func() {
		for {
			var mess string
			fmt.Println("Please enter message: ")
			fmt.Scan(&mess)
			message := &handlers.ChatWithUserMessage{
				Message:  mess,
				Receiver: receiver,
				Sender:   user.Id.String(),
			}

			if errReq := stream.Send(message); errReq != nil {
				log.Printf("Error sending message: %v", errReq)
				break
			}
		}
	}()

	go func() {
		for {
			resp, errRcv := stream.Recv()
			if errRcv == io.EOF {
				log.Println("End of file")
				break
			}

			if errRcv != nil {
				log.Printf("Error receiving message: %v", errRcv)

				break
			}

			log.Printf("message: %v", resp.GetMessage())
		}
	}()

	select {}
}
