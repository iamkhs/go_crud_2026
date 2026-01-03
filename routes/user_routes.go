package routes

import (
	"go_crud_2026/auth"
	"go_crud_2026/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	userRoutes.Use(auth.JWTMiddleware()) // protect all the /users route
	{
		userRoutes.GET("", handlers.GetUsers)
		userRoutes.GET("/:id", handlers.GetUserById)
		userRoutes.POST("", handlers.CreateUser)
		userRoutes.PUT("/:id", handlers.UpdateUser)
		userRoutes.DELETE("/:id", handlers.DeleteUser)
	}
}
