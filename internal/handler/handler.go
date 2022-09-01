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
			user.GET("/", h.getUser)
			user.PUT("/", h.updateUser)
			user.DELETE("/", h.deleteUser)

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
				beats.GET("/", h.getBeatsByToken)
				beats.POST("/", h.createBeat)
				beats.PUT("/:id", h.updateBeat)
				beats.DELETE("/:id", h.deleteBeat)
			}
		}

		beats := api.Group("/beats")
		{
			beats.GET("/", h.getAllBeats)
		}

		users := api.Group("/users")
		{
			beats.GET("/:id", h.getUsersBeats)
			users.GET("/", h.getAllUsers)
		}
	}

	return router
}
