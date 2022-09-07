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

			carts := accounts.Group("/carts")
			{
				carts.POST("/", h.addBeatToCart)
				carts.GET("/", h.getAllBeatsFromCart)
				carts.DELETE("/", h.deleteAllBeatsInCart)
				carts.DELETE("/:id", h.deleteBeatInCart)
			}

			playlists := accounts.Group("/playlists")
			{
				playlists.POST("/", h.createPlaylist)
				playlists.GET("/", h.getAllAccountsPlaylists)
				playlists.GET("/:id", h.getAccountsPlaylistById)
				playlists.PUT("/:id", h.updateAccountsPlaylist)
				playlists.DELETE("/:id", h.deleteAccountsPlaylist)
			}

			artists := accounts.Group("/artists")
			{
				artists.POST("/", h.createArtist)
				artists.GET("/", h.getArtistByToken)
				artists.DELETE("/:id", h.deleteArtist)

				beats := artists.Group("/beats")
				{
					beats.POST("/", h.createBeat)
					beats.GET("/", h.getAllArtistsBeats)
					beats.GET("/:id", h.getArtistsBeatById)
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
			beats.GET("/", h.getAllBeats)
			beats.GET("/:id", h.getBeatById)
		}

		playlists := api.Group("/playlists")
		{
			playlists.GET("/", h.getAllPlaylists)
			playlists.GET("/:id", h.getPlaylistById)
		}
	}

	return router
}
