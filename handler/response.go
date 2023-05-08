package handler

import (
	"fmt"

	"github.com/LucasRufo/golang-first-api/schemas"
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

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                      `json:"message"`
	Data    schemas.OpportunityResponse `json:"data"`
}
