package handlers

import (
	"go_crud_2026/models"
	"go_crud_2026/services"
	"go_crud_2026/utils"
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
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, _ := c.Get("userID")
	company.UserId = userId.(uint)
	service.CreateCompany(company)
	utils.SendSuccessResponse(c, http.StatusCreated, "Company created successfully", company)
}
