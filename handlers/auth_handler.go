package handlers

import (
	"go_crud_2026/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authService *services.UserService

func SetAuthService(service *services.UserService) {
	authService = service
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, ok := authService.Login(req.Email, req.Password)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
