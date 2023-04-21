package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendMessage(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.Trim(token, "\"")
	id, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var message app.Message
	err = json.NewDecoder(c.Request.Body).Decode(&message)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	message.Time = time.Now()
	message.Sender = user.Username

	err = h.services.SaveMessage(message)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) ReadMessages(c *gin.Context) {
	query_params := c.Request.URL.Query()
	username := query_params["user"][0]

	messages, err := h.services.ReadUserMessages(username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, struct {
		Messages []app.Message
	}{
		Messages: messages,
	})
}
