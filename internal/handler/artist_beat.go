package handler

import (
	"net/http"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createBeat(c *gin.Context) {
	artistId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input beatstore.Beat
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Beat.Create(artistId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllBeatsByToken(c *gin.Context) {
	artistId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	beats, err := h.service.Beat.GetAllArtistsBeats(artistId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, beats)
}

func (h *Handler) updateArtistsBeat(c *gin.Context) {

}

func (h *Handler) deleteArtistsBeat(c *gin.Context) {

}
