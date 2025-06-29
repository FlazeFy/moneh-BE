package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func AuthRouter(r *gin.Engine, redisClient *redis.Client, authController AuthController) {
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
