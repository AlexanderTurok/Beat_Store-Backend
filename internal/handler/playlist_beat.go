package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) addBeatToPlaylist(c *gin.Context) {
	playlistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "parameter playlist id is empty")
		return
	}

	beatId, err := strconv.Atoi(c.Param("beat_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "parameter beat id is empty")
		return
	}

	if err := h.service.Playlist.AddBeat(playlistId, beatId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handlers) getAllBeatsFromPlaylistByToken(c *gin.Context) {
	playlistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "parameter playlist id is empty")
		return
	}

	beats, err := h.service.Playlist.GetAllBeats(playlistId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, beats)
}

func (h *Handlers) deleteBeatFromPlaylist(c *gin.Context) {
	playlistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "parameter playlist id is empty")
		return
	}

	beatId, err := strconv.Atoi(c.Param("beat_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "parameter beat id is empty")
		return
	}

	if err := h.service.Playlist.DeleteBeat(playlistId, beatId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
