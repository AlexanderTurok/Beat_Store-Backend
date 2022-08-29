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

	user := router.Group("/user")
	{
		user.POST("/sign-up")
		user.POST("/sign-in")
		user.GET("/")
		user.GET("/:id")
		user.PUT("/:id")
		user.DELETE("/:id")

		cart := user.Group("/cart")
		{
			cart.POST("/")
			cart.GET("/")
			cart.GET("/:id")
			cart.PUT("/:id")
			cart.DELETE("/")
			cart.DELETE("/:id")
		}
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
