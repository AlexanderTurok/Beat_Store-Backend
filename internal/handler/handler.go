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

func (h *Handlers) InitRoutes() *gin.Engine {
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

			// payments := accounts.Group("/payments")
			// {

			// }

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

		confirmations := api.Group("/accounts/confirmations")
		{
			confirmations.GET("/:username", h.confirmAccount)
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
