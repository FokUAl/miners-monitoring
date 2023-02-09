package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/auth/sign-in")
	}

	id, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/auth/sign-in")
	} else {
		c.Set("id", id)
	}

}
