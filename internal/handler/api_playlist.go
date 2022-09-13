package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllAccountsPlaylists(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	playlists, err := h.service.Playlist.GetAllAccountsPlaylists(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, playlists)
}

func (h *Handler) getAllBeatsFromPlaylist(c *gin.Context) {

}
