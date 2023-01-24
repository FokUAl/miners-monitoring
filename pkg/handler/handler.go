package handler

import (
	"github.com/HgCl2/rock_scissors_paper/pkg/service"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/sign-up", h.getSignUp)
		auth.GET("/sign-in", h.getLogin)
		auth.POST("/google/login", h.googleLogin)

	}

	hall := router.Group("/hall", h.userIdentity)
	{
		hall.GET("/", h.hall)
		hall.POST("/create-room", h.createRoom)
		hall.GET("/gameroom", h.getGameRoom)
		hall.POST("/gameroom", h.postGameRoom)
	}

	// static routes
	router.Use(static.Serve("/", static.LocalFile("./ui", true)))

	return router
}
