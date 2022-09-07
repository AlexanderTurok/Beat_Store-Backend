package handler

import (
	"net/http"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAccountByToken(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	account, err := h.service.Account.Get(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *Handler) updateAccount(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var accountInput beatstore.AccountUpdateInput
	if err := c.BindJSON(&accountInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Account.Update(accountId, accountInput); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteAccount(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// if err := h.service.User.Delete(userId); err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, statusResponse{"ok"})
}
