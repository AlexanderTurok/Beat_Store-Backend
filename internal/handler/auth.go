package handler

import (
	"net/http"

	"github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) signUp(c *gin.Context) {
	var accountInput model.Account
	if err := c.BindJSON(&accountInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.Auth.CreateAccount(accountInput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handlers) signIn(c *gin.Context) {
	var accountInput signInInput
	if err := c.BindJSON(&accountInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Auth.GenerateToken(accountInput.Email, accountInput.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
