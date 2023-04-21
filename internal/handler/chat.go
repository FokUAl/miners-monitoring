package handler

import (
	"encoding/json"
	"fmt"
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
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
	message.SenderRole = user.Role

	err = h.services.SaveMessage(message)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) ReadMessages(c *gin.Context) {
	var username string

	err := json.NewDecoder(c.Request.Body).Decode(&username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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

func (h *Handler) GetSenders(c *gin.Context) {
	senders, err := h.services.GetSenders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("GetSenders: %s", err.Error()))
		return
	}

	type SenderStat struct {
		Name   string
		IsRead bool
	}
	var listOfSenders []SenderStat
	for _, sender := range senders {
		var temp SenderStat
		temp.Name = sender

		messages, err := h.services.ReadUserMessages(sender)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError,
				fmt.Sprintf("GetSenders: %s", err.Error()))
		}

		var isReadSender bool = true
		for _, msg := range messages {
			if !msg.IsRead {
				isReadSender = false
				break
			}
		}

		temp.IsRead = isReadSender

		listOfSenders = append(listOfSenders, temp)
	}

	c.JSON(http.StatusOK, struct {
		List []SenderStat
	}{
		List: listOfSenders,
	})
}
