package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getHome(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Printf("getHome: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	err = t.Execute(c.Writer, nil)
	if err != nil {
		log.Printf("getHome: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}
