package handler

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	Hub      *app.Hub
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/google/login", h.googleLogin)

	}

	home := router.Group("/", h.userIdentity)
	{
		home.GET("/home", h.getHome)

		home.POST("/add", h.addMiner)

		home.GET("/grid", h.minersGrid)

		home.GET("/asic", h.getMinerCharacteristics)

		home.GET("/find-asic-ip", h.FindDeviceIP)

		home.GET("/user-info", h.GetUserInfo)
		home.GET("/get-all-users", h.GetAllUsers)

		home.POST("/update-asic-info", h.UpdateAsicInfo)
		home.POST("/delete-device", h.DeleteDevice)
		home.POST("/comment-device", h.CommentDevice)
		home.POST("/delete-comment", h.DeleteComment)
		home.POST("/edit-comment", h.EditComment)

		home.GET("/read-user-messages", h.ReadUserMessages)
		home.POST("/read-messages-from", h.ReadMessages)
		home.POST("/send-message", h.SendMessage)

		home.GET("/get-senders", h.GetSenders)

		home.POST("/get-kernel-log", h.GetKernelLog)
	}

	return router
}
