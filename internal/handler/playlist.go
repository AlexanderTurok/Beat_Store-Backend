package handler

import (
	"net/http"
	"strconv"

	"github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) createPlaylist(c *gin.Context) {
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

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": playlistId,
	})
}

func (h *Handlers) getPlaylistByToken(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	playlistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid playlist id")
		return
	}

	playlist, err := h.service.Playlist.GetAccountsPlaylist(accountId, playlistId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, playlist)
}

func (h *Handlers) getAllPlaylistsByToken(c *gin.Context) {
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

func (h *Handlers) updatePlaylist(c *gin.Context) {
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

func (h *Handlers) deletePlaylist(c *gin.Context) {
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

func (h *Handlers) getPlaylistById(c *gin.Context) {
	playlistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid playlist id")
		return
	}

	playlist, err := h.service.Playlist.Get(playlistId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, playlist)

}

func (h *Handlers) getAllPlaylists(c *gin.Context) {
	playlists, err := h.service.Playlist.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, playlists)
}

func (h *Handlers) getAccountsPlaylist(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid account id")
		return
	}

	playlistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid playlist id")
		return
	}

	playlist, err := h.service.Playlist.GetAccountsPlaylist(accountId, playlistId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, playlist)
}

func (h *Handlers) getAllAccountsPlaylists(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid account id")
		return
	}

	playlists, err := h.service.Playlist.GetAllAccountsPlaylists(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, playlists)
}
