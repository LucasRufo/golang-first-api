package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func sendSuccess(c *gin.Context, code int, op string, data interface{}) {
	c.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull", op),
		"data":    data,
	})
}
