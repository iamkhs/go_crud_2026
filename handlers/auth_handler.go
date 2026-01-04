package handlers

import (
	"go_crud_2026/dto/request"
	"go_crud_2026/models"
	"go_crud_2026/services"
	"go_crud_2026/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authService *services.UserService

func SetAuthService(service *services.UserService) {
	authService = service
}

func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, ok := authService.Login(req.Email, req.Password)
	if !ok {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Login successful", gin.H{"token": token})
}

func Create(c *gin.Context) {
	var req request.UserCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		FullName: req.FullName,
		Email:    req.Email,
	}

	_, err := authService.CreateUser(user)
	if err != nil {
		if err.Error() == "email already exists" {
			utils.SendErrorResponse(c, http.StatusConflict, err.Error())
			return
		}
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to register user: "+err.Error())
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, "User created successfully, An Otp has been sent to your email", nil)
}

func VerifyOtp(c *gin.Context) {
	var req request.VerifyOtpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := authService.VerifyAndCompleteRegistration(req.Email, req.Otp, req.Password, req.CompanyName, req.EmployeeSize)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Account verified and registration completed successfully", nil)
}
