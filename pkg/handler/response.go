package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, code int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(code, Error{message})
}
