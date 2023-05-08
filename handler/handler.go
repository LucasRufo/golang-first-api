package handler

import (
	"fmt"
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
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing opportunities")
	}

	sendSuccess(c, http.StatusOK, "get-opportunities", opportunities)
}

func GetOpportunityByIdHandler(c *gin.Context) {
	id := c.Param("id")

	opportunity := schemas.Opportunity{}

	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	sendSuccess(c, http.StatusOK, "get-opportunity-by-id", opportunity)
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
	request := UpdateOpeningRequest{}

	c.BindJSON(&request)

	err := request.Validate()

	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schemas.Opportunity{}

	id := c.Param("id")

	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	if request.Role != "" {
		opportunity.Role = request.Role
	}

	if request.Company != "" {
		opportunity.Company = request.Company
	}

	if request.Location != "" {
		opportunity.Location = request.Location
	}

	if request.Remote != nil {
		opportunity.Remote = *request.Remote
	}

	if request.Link != "" {
		opportunity.Link = request.Link
	}

	if request.Salary > 0 {
		opportunity.Salary = request.Salary
	}

	if err := db.Save(&opportunity).Error; err != nil {
		logger.Errorf("error updating opportunity: %v", err.Error())
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error updating opportunity with id: %s", id))
		return
	}

	sendSuccess(c, http.StatusOK, "update-opportunity", opportunity)
}

func DeleteOpportunityHandler(c *gin.Context) {
	id := c.Param("id")

	opportunity := schemas.Opportunity{}

	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	if err := db.Delete(&opportunity).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting opportunity with id %s", id))
		return
	}

	sendSuccess(c, http.StatusNoContent, "delete-opportunity", opportunity)
}
