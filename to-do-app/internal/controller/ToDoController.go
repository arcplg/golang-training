package controller

import (
	"net/http"

	model "to-do-app/internal/models"
	db "to-do-app/pkg"

	"fmt"

	"github.com/gin-gonic/gin"
)

func AddTodoList(c *gin.Context) {
	var newToDo model.ToDo
	if err := c.ShouldBindJSON(&newToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Received Todo:", newToDo)

	_, err := db.DB.Exec("INSERT INTO to_do_list (title, description, is_completed) VALUES (?, ?, ?)", newToDo.Title, newToDo.Description, newToDo.IsCompleted)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	c.JSON(http.StatusCreated, newToDo)
}

func GetToDoList(c *gin.Context) {
	rows, err := db.DB.Query("Select * from to_do_list")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	defer rows.Close()

	var todoList []model.ToDo

	for rows.Next() {
		var todo model.ToDo
		rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)
		// if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error: %v", err)})
		// 	return
		// }

		todoList = append(todoList, todo)
	}

	c.JSON(http.StatusOK, todoList)
}
