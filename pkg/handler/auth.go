package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalimoldayev02/api-golang/models"
)

type AuthHandler struct {
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil { // передаем ссылку на объект/структуру BindJSON
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) signUp(c *gin.Context) {

}
