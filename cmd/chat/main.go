package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ngocthanh06/chatapp/cmd/chat/command"
	"github.com/ngocthanh06/chatapp/internal/handlers"
	"github.com/ngocthanh06/chatapp/internal/middleware"
	"github.com/ngocthanh06/chatapp/internal/services"
	"github.com/ngocthanh06/chatapp/internal/storage"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.

	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	done <- true
}

func recovered() {
	log.Println("Start recover")
	if r := recover(); r != nil {
		log.Printf("Error: %v\n", r)
	}
	log.Println("End recover")
}

func main() {
	defer recovered()

	// register commandline
	command.Execute()

	//server := server.NewServer()

	//go gracefulShutdown(server, done)

	// Do with http/ rest api
	//err := server.ListenAndServe()
	//if err != nil && err != http.ErrServerClosed {
	//	panic(fmt.Sprintf("http server error: %s", err))
	//}

	// connection database
	storage.New()

	// check connection is running
	lis, err := net.Listen("tcp", ":50070")

	if err != nil {
		log.Fatalf("err while create listener %v\n", err)
	}

	// Do with grpc
	gRpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.JWTAuthenticationInterceptor),
		//grpc.UnaryInterceptor(
		//	grpc_middleware.ChainUnaryServer(
		//		middleware.JWTAuthenticationInterceptor,
		//	),
		//),
		grpc.StreamInterceptor(middleware.JWTAuthenticationStreamInterceptor),
	)

	handlers.RegisterChatAppServiceServer(gRpcServer, &services.ChatAppServiceServer{
		Db:      storage.GetDB(),
		Clients: make(map[handlers.ChatAppService_ChatWithUserStreamServer]string),
		MsgChan: make(chan *handlers.ChatWithUserMessage),
	})

	go func() {
		log.Println("Start gRpc server port 50070")
		if errServer := gRpcServer.Serve(lis); errServer != nil {
			log.Fatalf("err while server %v\n", err)
		}

		fmt.Println("Connection gRpc is running...!")
	}()

	log.Println("Graceful shutdown complete")

	select {}
}
