package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf("ERROR %s\n", message)
	c.JSON(statusCode, struct {
		Code    int
		Message string
	}{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
	})
}
