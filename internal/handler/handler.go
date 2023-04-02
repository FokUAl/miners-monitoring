package handler

import (
	"github.com/FokUAl/miners-monitoring/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

type Notification struct {
	Message string
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	//go SaveMinerData(h)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/google/login", h.googleLogin)

	}

	home := router.Group("/", h.userIdentity)
	{
		home.GET("/home", h.getHome)
		home.POST("/home", h.getHome)

		home.POST("/add", h.addMiner)

		home.GET("/grid", h.minersGrid)

		home.GET("/asic", h.getMinerCharacteristics)

		home.GET("/find-asic-ip", h.FindDeviceIP)

		home.GET("/user-info", h.GetUserInfo)

		home.POST("/update-asic-info", h.UpdateAsicInfo)
		home.POST("/delete-device", h.DeleteDevice)
		home.POST("/comment-device", h.CommentDevice)
		home.POST("/delete-comment", h.DeleteComment)
		home.POST("/edit-comment", h.EditComment)
	}

	return router
}
