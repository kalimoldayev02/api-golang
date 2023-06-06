package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalimoldayev02/api-golang/models"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

  id, err := h.sevices.TodoList.Create(id, input)
  if err != nil {
  	newErrorResponse(c, http.)
  }
}
