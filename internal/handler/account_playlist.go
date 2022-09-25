package handler

import (
	"net/http"
	"strconv"

	"github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPlaylist(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.Playlist
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	playlistId, err := h.service.Playlist.Create(accountId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": playlistId,
	})
}

func (h *Handler) getAllPlaylistsByToken(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	playlists, err := h.service.Playlist.GetAllAccountsPlaylists(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, playlists)
}

func (h *Handler) updatePlaylist(c *gin.Context) {
	playlistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var updateInput model.PlaylistUpdateInput
	if err := c.BindJSON(&updateInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Playlist.Update(playlistId, updateInput); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteAccountsPlaylist(c *gin.Context) {
	playlistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.Playlist.Delete(playlistId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
