package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addBeat(c *gin.Context) {
	// if err := h.service.CreateBeat(ctx.Body); err != nil {
	// newErrorResponse(c)
	// }

	// c.JSON(http.StatusOK)
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

}

func (h *Handler) updateBeat(c *gin.Context) {

}

func (h *Handler) deleteBeat(c *gin.Context) {

}
