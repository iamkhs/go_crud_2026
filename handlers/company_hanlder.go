package handlers

import (
	"go_crud_2026/models"
	"go_crud_2026/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var service *services.CompanyService

func SetCompanyService(s *services.CompanyService) {
	service = s
}

func CreateCompany(c *gin.Context) {
	var company models.Company
	err := c.ShouldBindJSON(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userId, _ := c.Get("userID")
	company.UserId = userId.(uint)
	service.CreateCompany(company)
	c.JSON(http.StatusCreated, company)
}
