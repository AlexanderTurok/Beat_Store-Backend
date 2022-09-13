package handler

import (
	"net/http"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPlaylist(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input beatstore.Playlist
	if err := c.BindJSON(input); err != nil {
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

func (h *Handler) getAllAccountsPlaylists(c *gin.Context) {

}

func (h *Handler) getAccountsPlaylistById(c *gin.Context) {

}

func (h *Handler) updateAccountsPlaylist(c *gin.Context) {

}

func (h *Handler) deleteAccountsPlaylist(c *gin.Context) {

}
