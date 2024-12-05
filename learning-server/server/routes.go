package server

import (
	question "learning-server/api/controller"
	"learning-server/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
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
	router.GET("/graphql", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
	router.POST("/graphql", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	return router
}
