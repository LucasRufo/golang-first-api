package router

import (
	"github.com/LucasRufo/golang-first-api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	v1.GET("/opportunity", handler.GetOpportunityHandler)
	v1.GET("/opportunity/:id", handler.GetOpportunityByIdHandler)
	v1.POST("/opportunity", handler.CreateOpportunityHandler)
	v1.PUT("/opportunity", handler.UpdateOpportunityHandler)
	v1.DELETE("/opportunity", handler.DeleteOpportunityHandler)
}
