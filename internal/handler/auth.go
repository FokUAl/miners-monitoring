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

	userId, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, struct {
		UserID int
	}{
		UserID: userId,
	})
}

type signInInput struct {
	Username string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	}

	c.JSON(http.StatusOK, struct {
		Token string
	}{
		Token: token,
	})
}

func (h *Handler) logOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/auth/sign-in")
}
