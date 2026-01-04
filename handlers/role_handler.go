package handlers

import (
	"go_crud_2026/models"
	"go_crud_2026/services"
	"go_crud_2026/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var roleService *services.RoleService

func SetUpRoleService(r *services.RoleService) {
	roleService = r
}

func CreateRole(c *gin.Context) {
	var role models.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	roleService.CreateRole(role)
	utils.SendSuccessResponse(c, http.StatusCreated, "Role created successfully", role)
}

func GetAllRoles(c *gin.Context) {
	roles := roleService.GetAllRoles()
	utils.SendSuccessResponse(c, http.StatusOK, "Roles fetched successfully", roles)
}
