package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
