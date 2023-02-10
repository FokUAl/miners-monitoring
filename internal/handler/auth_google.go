package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/gin-gonic/gin"
)

func (h *Handler) googleLogin(c *gin.Context) {
	defer c.Request.Body.Close()

	// parse the GoogleJWT that was POSTed from the front-end
	type parameters struct {
		GoogleJWT *string
	}
	decoder := json.NewDecoder(c.Request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Validate the JWT is valid
	claims, err := h.services.ValidateGoogleJWT(*params.GoogleJWT)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, "Invalid google auth")
		return
	}

	user := app.User{
		Username: claims.FirstName + "_" + claims.LastName,
		Email:    claims.Email,
		Password: "password",
	}

	user_id, err := h.services.Authorization.CreateUser(user)
	if err != nil || user_id == 0 {
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("googleLogin: can't create user: %s", err.Error()))
	}

	// create a JWT for OUR app and give it back to the client for future requests
	tokenString, err := h.services.GenerateToken(claims.Email, "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")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Couldn't make authentication token")
		return
	}

	c.JSON(http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	})
}
