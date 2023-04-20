package handler

import (
	"encoding/json"
	"net/http"
	"time"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/gin-gonic/gin"
)

func (h *Handler) WriteMessage(c *gin.Context) {
	var message app.Message
	err := json.NewDecoder(c.Request.Body).Decode(&message)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	message.Time = time.Now()

	err = h.services.SaveMessage(message)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
