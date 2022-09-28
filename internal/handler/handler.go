package handler

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/service"
	"github.com/AlexanderTurok/beat-store-backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	service *service.Services
	manager auth.TokenManager
}

func NewHandlers(service *service.Services, manager auth.TokenManager) *Handlers {
	return &Handlers{
		service: service,
		manager: manager,
	}
}

func (h *Handlers) InitRouter() *gin.Engine {
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
			accounts.GET("/", h.getAccountByToken)
			accounts.PUT("/", h.updateAccount)
			accounts.DELETE("/", h.deleteAccount)

			products := accounts.Group("/products")
			{
				products.POST("/", h.buyProducts)
				products.GET("/", h.getAllBoughtProducts)
				products.GET("/:id", h.getBoughtProduct)
				products.DELETE("/:id", h.returnBoughtProduct)
			}

			playlists := accounts.Group("/playlists")
			{
				playlists.POST("/", h.createPlaylist)
				playlists.GET("/", h.getAllPlaylistsByToken)
				playlists.GET("/:id", h.getPlaylistByToken)
				playlists.PUT("/:id", h.updatePlaylist)
				playlists.DELETE("/:id", h.deletePlaylist)

				beats := playlists.Group(":id/beats")
				{
					beats.POST("/:beat_id", h.addBeatToPlaylist)
					beats.GET("/", h.getAllBeatsFromPlaylist)
					beats.GET("/:beat_id", h.getBeatFromPlaylist)
				}
			}

			artists := api.Group("/artists")
			{
				artists.POST("/", h.createArtist)
				artists.GET("/", h.getArtistByToken)
				artists.PUT("/", h.updateArtist)
				artists.DELETE("/", h.deleteArtist)

				products := artists.Group("/products")
				{
					products.GET("/", h.getAllProductsByToken)
					products.GET("/:id", h.getProductByToken)
					products.PUT("/:id", h.updateProduct)
					products.DELETE(":id", h.deleteProduct)
				}

				beats := artists.Group("/beats")
				{
					beats.POST("/", h.createBeat)
					beats.GET("/", h.getAllBeatsByToken)
					beats.GET("/:id", h.getBeatByToken)
					beats.PUT("/:id", h.updateBeat)
					beats.DELETE("/:id", deleteBeat)
				}
			}
		}

		artists := api.Group("/artists")
		{
			artists.GET("/", h.getAllArtists)
			artists.GET("/", h.getArtistById)
		}

		playlists := api.Group("/playlists")
		{
			playlists.GET("/", h.getAllPlaylists)
			playlists.GET("/:id", h.getPlaylistById)
		}

		beats := api.Group("/beats")
		{
			beats.GET("/", h.getAllBeats)
			beats.GET("/:id", h.getBeatById)
		}
	}

	return router
}
