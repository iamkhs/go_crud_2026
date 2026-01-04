package main

import (
	"go_crud_2026/database"
	"go_crud_2026/handlers"
	"go_crud_2026/repositories"
	"go_crud_2026/routes"
	"go_crud_2026/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect database
	database.Connect()

	// Setup repository & service
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)

	// Inject service into handlers
	handlers.SetUserService(userService)
	handlers.SetAuthService(userService)

	// Register routes
	routes.RegisterUserRoutes(r)

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Create)

	_ = r.Run(":8080")
}
