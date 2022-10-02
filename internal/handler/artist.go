package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) createArtist(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.Artist.Create(accountId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	url, err := h.service.Payment.CreatePaymentAccount(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusFound, url)
}

func (h *Handlers) getArtistById(c *gin.Context) {
	artistId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid artist id")
		return
	}

	artist, err := h.service.Artist.Get(artistId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, artist)
}

func (h *Handlers) getAllArtists(c *gin.Context) {
	artists, err := h.service.Artist.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, artists)
}

func (h *Handlers) getArtistByToken(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	artist, err := h.service.Artist.Get(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, artist)
}

type UpdateArtist struct {
	StripeId string
}

func (h *Handlers) updateArtist(c *gin.Context) {

}

type ArtistPassword struct {
	Password string `json:"password" db:"password_hash"`
}

func (h *Handlers) deleteArtist(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var password ArtistPassword
	if err := c.BindJSON(&password); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Artist.Delete(accountId, password.Password); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
