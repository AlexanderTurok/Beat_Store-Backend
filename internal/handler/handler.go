package handler

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/service"
	"github.com/AlexanderTurok/beat-store-backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	manager auth.TokenManager
}

func NewHandler(service *service.Service, manager auth.TokenManager) *Handler {
	return &Handler{
		service: service,
		manager: manager,
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
			accounts.GET("/", h.getAccountByToken)
			accounts.PUT("/", h.updateAccount)
			accounts.DELETE("/", h.deleteAccount)

			beats := accounts.Group("/beats")
			{
				beats.POST("/", h.buyBeat)
				beats.GET("/", h.getAllBoughtBeats)
				beats.GET("/:id", h.getBoughtBeatById)
				beats.DELETE("/:id", h.returnBoughtBeat)
			}

			playlists := accounts.Group("/playlists")
			{
				playlists.POST("/", h.createPlaylist)
				playlists.GET("/", h.getAllPlaylistsByToken)
				playlists.PUT("/:id", h.updatePlaylist)
				playlists.DELETE("/:id", h.deleteAccountsPlaylist)

				beats := playlists.Group(":id/beats")
				{
					beats.POST("/:beat_id", h.addBeatToPlaylist)
					beats.GET("/", h.getAllBeatsFromPlaylistByToken)
					beats.DELETE("/:beat_id", h.deleteBeatFromPlaylist)
				}
			}

			artists := accounts.Group("/artists")
			{
				artists.POST("/", h.createArtist)
				artists.GET("/", h.getArtistByToken)
				artists.DELETE("/", h.deleteArtist)

				beats := artists.Group("/beats")
				{
					beats.POST("/", h.createBeat)
					beats.GET("/", h.getAllBeatsByToken)
					beats.PUT("/:id", h.updateArtistsBeat)
					beats.DELETE("/:id", h.deleteArtistsBeat)
				}
			}
		}

		artists := api.Group("/artists")
		{
			artists.GET("/", h.getAllArtists)
			artists.GET("/:id", h.getArtistById)
		}

		beats := api.Group("/beats")
		{
			beats.POST("/:id", h.getBeatById)
			beats.GET("/", h.getAllBeats)
			beats.GET("/:id", h.getAllArtistsBeats)
		}

		playlists := api.Group("/playlists")
		{
			playlists.GET("/:id", h.getAllAccountsPlaylists)
			playlists.GET("/:id/beats", h.getAllBeatsFromPlaylist)
		}
	}

	return router
}
