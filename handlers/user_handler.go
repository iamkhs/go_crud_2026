package handlers

import (
	"go_crud_2026/models"
	"go_crud_2026/services"
	"go_crud_2026/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userService *services.UserService

func SetUserService(service *services.UserService) {
	userService = service
}

func GetUsers(c *gin.Context) {
	users := userService.GetAllUsers()
	utils.SendSuccessResponse(c, http.StatusOK, "Users retrieved successfully", users)
}

func GetUserById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, found := userService.GetById(id)

	if !found {
		utils.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := userService.UpdateUser(id, updatedUser)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "User not found or update failed")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "User updated successfully", user)
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	isDelete := userService.DeleteUser(id)
	if !isDelete {
		utils.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}
