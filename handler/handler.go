package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "get",
	})
}

func GetOpportunityByIdHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "get by id",
	})
}

func CreateOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "post",
	})
}

func UpdateOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "put",
	})
}

func DeleteOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete",
	})
}
