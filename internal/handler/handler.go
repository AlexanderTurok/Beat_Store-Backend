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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		user := api.Group("/user", h.userIdentity)
		{
			user.GET("/", h.getAllUsers)
			user.GET("/:id", h.getUserById)
			user.PUT("/:id", h.updateUser)
			user.DELETE("/:id", h.deleteUser)

			cart := user.Group("/cart")
			{
				cart.POST("/", h.addBeatToCart)
				cart.GET("/", h.getAllBeatsFromCart)
				cart.GET("/:id", h.getBeatByIdFromCart)
				cart.DELETE("/", h.deleteAllBeatsInCart)
				cart.DELETE("/:id", h.deleteBeatInCart)
			}

			beats := user.Group("/beats")
			{
				beats.POST("/", h.createBeat)
				beats.PUT("/:id", h.updateBeat)
				beats.DELETE("/:id", h.deleteBeat)
			}
		}

		beats := api.Group("/beats")
		{
			beats.GET("/", h.getAllBeats)
		}
	}

	return router
}
