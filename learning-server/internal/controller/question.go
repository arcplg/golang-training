package question

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct{}

func NewHandler() *QuestionHandler {
	return &QuestionHandler{}
}

func (h *QuestionHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List API - List"})
}

func (h *QuestionHandler) Detail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Detail API - Detail"})
}
