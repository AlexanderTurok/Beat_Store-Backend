package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAccount(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// user, err := h.service.User.Get(userId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, user)
}

func (h *Handler) getAllAccounts(c *gin.Context) {
	// users, err := h.service.User.GetAll()
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, users)
}

func (h *Handler) updateAccount(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// var userInput beatstore.AccountUpdateInput
	// if err := c.BindJSON(&userInput); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// if err := h.service.User.Update(userId, userInput); err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, statusResponse{"ok"})
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
