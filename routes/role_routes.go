package routes

import (
	"go_crud_2026/auth"
	"go_crud_2026/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoleRoutes(rg *gin.RouterGroup) {
	roleRoutes := rg.Group("/roles")

	roleRoutes.Use(auth.JWTMiddleware())
	{
		roleRoutes.GET("", handlers.GetAllRoles)
		roleRoutes.POST("", handlers.CreateRole)
	}
}
