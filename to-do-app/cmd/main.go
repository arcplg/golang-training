package main

import (
	"fmt"
	"log"
	db "to-do-app/pkg"

	"to-do-app/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// Kết nối MySQL
	_, err := db.InitDB()
	// check have error
	if err != nil {
		fmt.Println("Connected to MySQL ERROR!")
		log.Fatal(err)
	}
	defer db.DB.Close()

	fmt.Println("Connected to MySQL!")

	db.MigrateUserTable(db.DB)
	db.MigrateToDoListTable(db.DB)

	router := gin.Default()

	router.POST("/login", controller.Login)

	router.POST("/to-do-list", controller.AddTodoList)
	router.GET("/to-do-list", controller.GetToDoList)

	router.POST("/user/register", controller.AddUser)
	router.GET("/user-list", controller.GetUserList)

	router.Run("localhost:8088")
}
