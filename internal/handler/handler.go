package handler

import (
	"net/http"

	"github.com/FokUAl/miners-monitoring/internal/service"
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

	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/google/login", h.googleLogin)
		auth.GET("/logout", h.logOut)

	}

	home := router.Group("/", h.userIdentity)
	{
		home.GET("/home", h.getHome)
		home.POST("/home", h.getHome)

		home.GET("/add", h.getAddMiner)
		home.POST("/add", h.addMiner)

		home.GET("/grid", h.minersGrid)

		home.GET("/asic", h.getMinerCharacteristics)

		//home.GET("/find-asic-ip", h.GetFindDeviceIP)
		home.POST("/find-asic-ip", h.FindDeviceIP)
	}

	// static routes
	router.StaticFS("/static/", http.Dir("./ui/static"))

	return router
}
