package server

import (
	"fmt"
	"github.com/ngocthanh06/chatapp/internal/storage"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port int
	db   storage.DatabaseService
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	NewServer := &Server{
		port: port,
		db:   storage.New(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}
