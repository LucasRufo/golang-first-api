package router

import (
	"github.com/LucasRufo/golang-first-api/docs"
	"github.com/LucasRufo/golang-first-api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	handler.Init()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := r.Group(basePath)

	v1.GET("/opportunity", handler.GetOpportunityHandler)
	v1.GET("/opportunity/:id", handler.GetOpportunityByIdHandler)
	v1.POST("/opportunity", handler.CreateOpportunityHandler)
	v1.PUT("/opportunity/:id", handler.UpdateOpportunityHandler)
	v1.DELETE("/opportunity/:id", handler.DeleteOpportunityHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
