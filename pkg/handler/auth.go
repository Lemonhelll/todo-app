package handler

import (
	"github.com/gin-gonic/gin"
	myTodo_app "myTodo-app"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input myTodo_app.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
