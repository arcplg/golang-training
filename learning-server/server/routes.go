package server

import (
	"bytes"
	"encoding/json"
	"io"
	question "learning-server/api/controller"
	"learning-server/graph"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// REST routes
	api := router.Group("/api")

	questionHandler := question.NewHandler()
	api.GET("/question", questionHandler.List)
	api.GET("/question/:id", questionHandler.Detail)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(transport.SSE{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	router.Any("/graphql", func(c *gin.Context) {
		appEnv := os.Getenv("APP_ENV")
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot read request body"})
			return
		}
		if appEnv == "prod" {
			log.Println("Raw request body:", string(body))
		} else if len(body) != 0 {
			var jsonData interface{}
			if err := json.Unmarshal(body, &jsonData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
				return
			}
			jsonPretty, _ := json.MarshalIndent(jsonData, "", "  ")
			log.Println("Raw request body Parsed JSON:", string(jsonPretty))
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		srv.ServeHTTP(c.Writer, c.Request)
	})

	return router
}
