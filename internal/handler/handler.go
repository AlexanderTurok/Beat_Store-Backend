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
		accounts := api.Group("/accounts", h.userIdentity)
		{
			accounts.POST("/")
			accounts.GET("/")
			accounts.GET("/:id")
			accounts.PUT("/:id")
			accounts.DELETE("/:id")

			carts := accounts.Group("/carts")
			{
				carts.POST("/")
				carts.GET("/")
				carts.DELETE("/:")
				carts.DELETE("/:id")
			}

			artists := accounts.Group("/artists")
			{
				artists.GET("/")
				artists.GET("/:id")
				artists.DELETE("/:id")

				beats := accounts.Group("/beats")
				{
					beats.POST("/")
					beats.GET("/")
					beats.GET("/:id")
					beats.PUT("/:id")
					beats.DELETE("/:id")
				}

				playlists := accounts.Group("/playlists")
				{
					playlists.POST("/")
					playlists.GET("/")
					playlists.GET("/:id")
					playlists.PUT("/:id")
					playlists.DELETE("/:id")
				}
			}
		}

		beats := api.Group("/beats")
		{
			beats.GET("/")
			beats.GET("/:id")
		}

		playlists := api.Group("/playlists")
		{
			playlists.GET("/")
			playlists.GET("/:id")
		}
	}

	return router
}
