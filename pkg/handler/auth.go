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

	h.ser
}

func (h *Handler) signIn(c *gin.Context) {

}
