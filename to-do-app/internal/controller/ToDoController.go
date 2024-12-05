package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	model "to-do-app/internal/models"
	db "to-do-app/pkg"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"to-do-app/internal/repository"
)

type ToDoController struct {
	repo repository.ToDoRepository
}

func NewToDoController(repo repository.ToDoRepository) *ToDoController {
	return &ToDoController{repo: repo}
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(c *gin.Context) {
	fmt.Println("Headers:", c.Request.Header)

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true
	fmt.Println("New WebSocket connection established")

	for {
		var msg []byte
		if _, msg, err = ws.ReadMessage(); err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func (ctrl *ToDoController) AddTodoList(c *gin.Context) {
	var newToDo model.ToDo
	if err := c.ShouldBindJSON(&newToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Received Todo:", newToDo)

	result, err := db.DB.Exec("INSERT INTO to_do_list (title, description, is_completed, user_id) VALUES (?, ?, ?, 1)", newToDo.Title, newToDo.Description, newToDo.IsCompleted)

	if err != nil {
		fmt.Println("Error:", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve last insert ID"})
		return
	}

	fmt.Println("last id = ", lastInsertID)
	// Query the database for the newly inserted row
	var insertedToDo model.ToDo
	var createdAt []byte

	var updatedAt []byte

	row := db.DB.QueryRow("SELECT id, title, description, is_completed, user_id, created_at, updated_at FROM to_do_list WHERE id = ?", lastInsertID)
	err = row.Scan(&insertedToDo.ID, &insertedToDo.Title, &insertedToDo.Description, &insertedToDo.IsCompleted, &insertedToDo.UserId, &createdAt, &updatedAt)

	if err != nil {
		fmt.Println("Error fetching inserted row:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inserted data"})
		return
	}

	// Convert byte slice to time.Time
	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt)) // Adjust format based on your DB format
	if err != nil {
		fmt.Println("Error parsing time:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse time"})
		return
	}
	// Convert byte slice to time.Time
	parsedTime2, err := time.Parse("2006-01-02 15:04:05", string(updatedAt)) // Adjust format based on your DB format
	if err != nil {
		fmt.Println("Error parsing time:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse time"})
		return
	}

	insertedToDo.CreatedAt = parsedTime
	insertedToDo.UpdatedAt = parsedTime2

	// Send the new To-Do data to connected WebSocket clients
	toDoData, err := json.Marshal(gin.H{
		"action": "new_todo",
		"data":   insertedToDo,
	})

	broadcast <- toDoData // Send JSON data to WebSocket clients

	c.JSON(http.StatusCreated, insertedToDo)
}

func (ctrl *ToDoController) GetToDoList(c *gin.Context) {
	todos, err := ctrl.repo.GetAll()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (ctrl *ToDoController) GetToDoDetail(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id) // convert string to int
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	todo, err := ctrl.repo.GetDetailById(intID)

	if err != nil {
		fmt.Println("Error: ", err)

		c.JSON(http.StatusNotFound, gin.H{"error": "To-Do item not found"})
		return
	}

	// Return the To-Do item as JSON
	c.JSON(http.StatusOK, todo)
}

func (ctrl *ToDoController) EditToDoList(c *gin.Context) {
	//
	id := c.Param("id")
	intID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var toDo model.ToDo
	if err := c.ShouldBindJSON(&toDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	toDo.ID = intID
	if err := ctrl.repo.Update(&toDo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ToDo"})
		return
	}

	// insertedToDo, err := ctrl.repo.GetDetailById(intID)

	// toDoData, err := json.Marshal(gin.H{
	// 	"action": "new_todo",
	// 	"data":   insertedToDo,
	// })

	// broadcast <- toDoData // Send JSON data to WebSocket clients

	// toDo.ID, _ = strconv.Atoi(id) // Đảm bảo ID đúng kiểu
	data, _ := json.Marshal(toDo)
	broadcast <- data // Send JSON data to WebSocket clients

	c.JSON(http.StatusCreated, toDo)

	return
}

func (ctrl *ToDoController) DeleteToDoList(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := ctrl.repo.Delete(intID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ToDo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ToDo delete successfully"})
	return
}
