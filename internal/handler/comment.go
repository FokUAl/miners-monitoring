package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CommentDevice(c *gin.Context) {
	type inputCommentData struct {
		Content string `json:"comment"`
		Address string `json:"IP"`
	}
	var input inputCommentData
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("CommentDevice: %s", err.Error()))
		return
	}

	id := c.MustGet(userCtx).(int)
	user, err := h.services.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}

	err = h.services.Comment.Comment(input.Address, user.Username, input.Content)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}
}

func (h *Handler) DeleteComment(c *gin.Context) {
	var input app.Comment
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("DeleteComment: %s", err.Error()))
		return
	}

	err = h.services.Comment.DeleteComment(input.CreationDate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("DeleteComment: %s", err.Error()))
		return
	}
}

func (h *Handler) EditComment(c *gin.Context) {
	var input app.Comment
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("DeleteComment: %s", err.Error()))
		return
	}

	err = h.services.Comment.EditComment(input.CreationDate, input.Content)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("DeleteComment: %s", err.Error()))
		return
	}
}
