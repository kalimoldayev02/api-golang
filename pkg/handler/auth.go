package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalimoldayev02/api-golang/models"
)

type AuthHandler struct {
}

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil { // передаем ссылку на объект/структуру BindJSON
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.sevices.Auth.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.UserSignIn

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.sevices.Auth.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
