package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB) {
	// Auth Module
	userRepo := NewUserRepository(db)
	authService := NewAuthService(userRepo)
	authController := NewAuthController(authService)

	api := r.Group("/api/v1")
	{
		// Public Routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
	}
}
