package handler

import (
	"html/template"
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
	t, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		log.Printf("errorPage: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	errorText, err := c.Cookie("ErrorCode")
	// log.Println(errorText)
	err = t.Execute(c.Writer, errorText)
	if err != nil {
		log.Printf("errorPage: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}
