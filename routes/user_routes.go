package routes

import (
	"go_crud_2026/auth"
	"go_crud_2026/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoutes := rg.Group("/users")

	userRoutes.Use(auth.JWTMiddleware())
	{
		userRoutes.GET("", handlers.GetUsers)
		userRoutes.GET("/:id", handlers.GetUserById)
		userRoutes.PUT("/:id", handlers.UpdateUser)
		userRoutes.DELETE("/:id", handlers.DeleteUser)
	}
}
