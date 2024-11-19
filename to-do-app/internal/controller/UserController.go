package controller

import (
	"net/http"

	model "to-do-app/internal/models"
	db "to-do-app/pkg"

	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Received User:", newUser)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashedPassword)

	_, err := db.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", newUser.UserName, newUser.Email, newUser.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func GetUserList(c *gin.Context) {
	rows, err := db.DB.Query("Select * from users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	defer rows.Close()

	var userList []model.User

	for rows.Next() {
		var user model.User
		rows.Scan(&user.ID, &user.UserName, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		userList = append(userList, user)
	}

	c.JSON(http.StatusOK, userList)
}
