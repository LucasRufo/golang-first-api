package handler

import (
	"net/http"

	"github.com/LucasRufo/golang-first-api/config"
	"github.com/LucasRufo/golang-first-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger()
	db = config.GetDb()
}

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
	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	err := request.Validate()

	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schemas.Opportunity{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Errorf("error creating opportunity: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error creating opportunity")
		return
	}

	sendSuccess(c, http.StatusOK, "create-opening", opportunity)
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
