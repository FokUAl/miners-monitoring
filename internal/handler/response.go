package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf("ERROR %s\n", message)
	c.SetCookie("ErrorCode", strconv.Itoa(statusCode), 20, "/error", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/error")
}

func (h *Handler) errorPage(c *gin.Context) {
	errorText, err := c.Cookie("ErrorCode")
	if err != nil {
		log.Printf("errorPage: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

	errorCode, err := strconv.Atoi(errorText)
	if err != nil {
		log.Printf("errorPage: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
	c.JSON(errorCode, struct {
		Code int
	}{
		Code: errorCode,
	})
}
