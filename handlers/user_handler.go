package handlers

import (
	"go_crud_2026/services"
	"net/http"
	"strconv"

	"go_crud_2026/models"

	"github.com/gin-gonic/gin"
)

var userService *services.UserService

func SetUserService(service *services.UserService) {
	userService = service
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userService.GetAllUsers())
}

func GetUserById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, found := userService.GetById(id)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdUser, err := userService.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, found := userService.UpdateUser(id, updatedUser)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	isDelete := userService.DeleteUser(id)
	if !isDelete {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
