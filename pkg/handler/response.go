package handler

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	//log.Fatalln(message)
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}
