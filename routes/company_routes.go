package routes

import (
	"go_crud_2026/auth"
	"go_crud_2026/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCompanyRoutes(rg *gin.RouterGroup) {
	companyRoutes := rg.Group("/users/company")

	companyRoutes.Use(auth.JWTMiddleware())
	{
		companyRoutes.POST("", handlers.CreateCompany)
		companyRoutes.PUT("/:id", handlers.UpdateUser)
	}
}
