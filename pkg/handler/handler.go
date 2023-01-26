package handler

import (
	"net/http"

	"github.com/FokUAl/miners-monitoring/pkg/service"
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

	home := router.Group("/", h.userIdentity)
	{
		home.GET("/", h.getHome)
		home.POST("/", h.getHome)
		home.GET("/add", h.getAddMiner)
		home.POST("/add", h.addMiner)
		home.GET("/grid", h.minersGrid)
	}

	// static routes
	router.StaticFS("/static/", http.Dir("./ui/static"))

	return router
}
