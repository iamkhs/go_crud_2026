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
	emailService := services.NewEmailService()
	userService := services.NewUserService(userRepo, emailService)
	companyRepo := repositories.NewCompanyRepository()
	companyService := services.NewCompanyService(companyRepo)

	// Inject service into handlers
	handlers.SetUserService(userService)
	handlers.SetAuthService(userService)
	handlers.SetCompanyService(companyService)

	// Register routes
	v1 := r.Group("/api/v1")
	routes.RegisterUserRoutes(v1)
	routes.RegisterCompanyRoutes(v1)

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Create)

	_ = r.Run(":8080")
}
