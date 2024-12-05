package main

import (
	"fmt"
	"log"

	"to-do-app/internal/controller"
	"to-do-app/internal/repository"
	db "to-do-app/pkg"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MySQL
	_, err := db.InitDB()
	if err != nil {
		fmt.Println("Connected to MySQL ERROR!")
		log.Fatal(err)
	}
	defer db.DB.Close()

	fmt.Println("Connected to MySQL!")

	// Migrate database tables
	db.MigrateUserTable(db.DB)
	db.MigrateToDoListTable(db.DB)

	// Initialize Gin router
	router := gin.Default()

	// Apply CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	todoRepo := repository.NewToDoRepository(db.DB)
	todoCtrl := controller.NewToDoController(todoRepo)

	// Define routes
	router.POST("/api/to-do-list", todoCtrl.AddTodoList)
	router.GET("/api/to-do-list", todoCtrl.GetToDoList)
	router.GET("/api/to-do-list/:id", todoCtrl.GetToDoDetail)
	router.PUT("/api/to-do-list/:id", todoCtrl.EditToDoList)
	router.DELETE("/api/to-do-list/:id", todoCtrl.DeleteToDoList)

	router.POST("/api/login", controller.Login)
	router.POST("/api/user/register", controller.AddUser)
	router.GET("/api/user-list", controller.GetUserList)

	router.GET("/ws", controller.HandleConnections)

	go controller.HandleMessages()

	// Start the server
	router.Run("localhost:8088")
}
