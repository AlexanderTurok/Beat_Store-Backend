package handler

import (
	"net/http"
	"strconv"

	"github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) createBeat(c *gin.Context) {
	artistId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.Beat
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Beat.Create(artistId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handlers) getBeatByToken(c *gin.Context) {

}

func (h *Handlers) getAllBeats(c *gin.Context) {
	beats, err := h.service.Beat.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, beats)
}

func (h *Handlers) getBeatById(c *gin.Context) {
	beatId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	beat, err := h.service.Beat.Get(beatId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, beat)
}

func (h *Handlers) getAllBeatsByToken(c *gin.Context) {
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

func (h *Handlers) updateBeat(c *gin.Context) {
	beatId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input model.BeatUpdateInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Beat.Update(beatId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handlers) deleteBeat(c *gin.Context) {
	beatId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.Beat.Delete(beatId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handlers) getBeatFromPlaylist(c *gin.Context) {

}

func (h *Handlers) getAllBeatsFromPlaylist(c *gin.Context) {
	playlistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	beats, err := h.service.Playlist.GetAllBeats(playlistId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, beats)
}

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
