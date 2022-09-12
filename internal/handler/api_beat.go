package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllBeats(c *gin.Context) {
	beats, err := h.service.Beat.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, beats)
}

func (h *Handler) getAllArtistsBeats(c *gin.Context) {
	artistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	beats, err := h.service.Beat.GetAllArtistsBeats(artistId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, beats)
}

// func (h *Handler) getBeatById(c *gin.Context) {
// 	beatId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "invalid id")
// 		return
// 	}

// 	beat, err := h.service.Beat.Get(beatId)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, beat)
// }
