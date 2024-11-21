package question

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User API - GetQuestion"})
}

func (h *UserHandler) Register(router *gin.RouterGroup) {
	userGroup := router.Group("/question")
	userGroup.GET("/:id", h.GetQuestions)
}
