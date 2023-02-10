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
	c.Set("ErrorCode", statusCode)
	c.Redirect(http.StatusSeeOther, "/error")
}

func (h *Handler) errorPage(c *gin.Context) {
	errorCode := c.MustGet("ErrorCode").(int)

	c.JSON(errorCode, struct {
		Code int
	}{
		Code: errorCode,
	})
}
