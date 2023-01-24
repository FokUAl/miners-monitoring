package handler

import (
	"html/template"
	"log"
	"net/http"

	app "github.com/HgCl2/rock_scissors_paper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input app.Player

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getLogin(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/sign-in.html")
	if err != nil {
		log.Printf("getSignUp: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	err = t.Execute(c.Writer, nil)
	if err != nil {
		log.Printf("getSignUp: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

func (h *Handler) getSignUp(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/sign-up.html")
	if err != nil {
		log.Printf("getSignUp: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	err = t.Execute(c.Writer, nil)
	if err != nil {
		log.Printf("getSignUp: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
