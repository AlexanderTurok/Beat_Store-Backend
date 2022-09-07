package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) createBeat(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// var userInput beatstore.Beat
	// if err := c.BindJSON(&userInput); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// id, err := h.service.Beat.Create(userId, userInput)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"id": id,
	// })
}

func (h *Handler) getAllArtistsBeats(c *gin.Context) {

}

func (h *Handler) getArtistsBeatById(c *gin.Context) {

}

func (h *Handler) updateArtistsBeat(c *gin.Context) {

}

func (h *Handler) deleteArtistsBeat(c *gin.Context) {

}
