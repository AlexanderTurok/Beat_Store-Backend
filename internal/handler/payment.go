package handler

import (
	"net/http"

	"github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPaymentIntent(c *gin.Context) {
	var input model.PaymentInfo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	paymentIntent, err := h.service.Payment.CreatePaymentIntent(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, paymentIntent)
}
