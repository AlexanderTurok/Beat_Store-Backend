package handler

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/sign-up")
		api.POST("/sign-in")
	}

	beats := router.Group("/beats")
	{
		beats.POST("/")
		beats.GET("/")
		beats.GET("/:id")
		beats.PUT("/:id")
		beats.DELETE("/:id")
	}

	return router
}
