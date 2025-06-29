package auth

import (
	"moneh/modules/admin"
	"moneh/modules/user"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	// Auth Module
	userRepo := user.NewUserRepository(db)
	adminRepo := admin.NewAdminRepository(db)
	authService := NewAuthService(userRepo, adminRepo, redisClient)
	authController := NewAuthController(authService)

	api := r.Group("/api/v1")
	{
		// Public Routes
		auth := api.Group("/auths")
		{
			auth.POST("/register", authController.BasicRegister)
			auth.POST("/login", authController.BasicLogin)
			auth.POST("/signout", authController.BasicSignOut)
		}
	}
}
