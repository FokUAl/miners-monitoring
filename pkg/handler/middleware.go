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
		// newErrorResponse(c, http.StatusUnauthorized, err.Error())
		// return
	}

	userId, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/auth/sign-in")
		// newErrorResponse(c, http.StatusUnauthorized, err.Error())
		// return
	}

	c.Set(userCtx, userId)
}
