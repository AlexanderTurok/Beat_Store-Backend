package handler

import (
	"net/http"
	"strconv"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addBeat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var userInput beatstore.Beat
	if err := c.BindJSON(&userInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Beat.Create(userId, userInput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getBeatById(c *gin.Context) {
	beatId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	beat, err := h.service.Beat.GetById(beatId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, beat)
}

func (h *Handler) getAllBeats(c *gin.Context) {
	beats, err := h.service.Beat.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, beats)
}

func (h *Handler) updateBeat(c *gin.Context) {

}

func (h *Handler) deleteBeat(c *gin.Context) {

}
