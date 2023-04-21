package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.services.ListOfUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
	}

	c.JSON(200, struct {
		AllUsers []string
	}{
		AllUsers: users,
	})
}
