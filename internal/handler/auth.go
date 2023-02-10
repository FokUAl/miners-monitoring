package handler

import (
	"encoding/json"
	"net/http"

	app "github.com/FokUAl/miners-monitoring"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input app.User

	type SignUpInfo struct {
		Nickname string
		Password string
	}

	var info SignUpInfo
	err := json.NewDecoder(c.Request.Body).Decode(&info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.Username = info.Nickname
	input.Password = info.Password

	_, err = h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/auth/sign-in")
}

// empty
// func (h *Handler) getLogin(c *gin.Context) {
// 	t, err := template.ParseFiles("./ui/html/sign-in.html")
// 	if err != nil {
// 		log.Printf("getSignUp: %s\n", err.Error())
// 		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
// 			http.StatusInternalServerError)
// 		return
// 	}
// 	err = t.Execute(c.Writer, nil)
// 	if err != nil {
// 		log.Printf("getSignUp: %s\n", err.Error())
// 		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
// 			http.StatusInternalServerError)
// 	}
// }

// empty
// func (h *Handler) getSignUp(c *gin.Context) {
// 	t, err := template.ParseFiles("./ui/html/sign-up.html")
// 	if err != nil {
// 		log.Printf("getSignUp: %s\n", err.Error())
// 		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
// 			http.StatusInternalServerError)
// 		return
// 	}
// 	err = t.Execute(c.Writer, nil)
// 	if err != nil {
// 		log.Printf("getSignUp: %s\n", err.Error())
// 		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
// 			http.StatusInternalServerError)
// 	}
// }

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	input.Username = c.PostForm("nickname")
	input.Password = c.PostForm("password")

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	} else {
		c.SetCookie("token", token, 10000, "/", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func (h *Handler) logOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/auth/sign-in")
}
