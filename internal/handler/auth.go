package handler

import (
	"encoding/json"
	"net/http"

	app "github.com/FokUAl/miners-monitoring"

	"github.com/gin-gonic/gin"
)

type SignUpInfo struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input app.User

	var info SignUpInfo
	err := json.NewDecoder(c.Request.Body).Decode(&info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.Username = info.Nickname
	input.Password = info.Password
	input.Email = info.Email

	_, err = h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/auth/sign-in")
}

type signInInput struct {
	Username string `json:"user" binding:"required"`
	Password string `json:"pwd" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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
