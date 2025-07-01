package user

import (
	"moneh/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func UserRouter(r *gin.Engine, userController UserController, redisClient *redis.Client) {
	api := r.Group("/api/v1")
	{
		// Private Routes
		protected_user := api.Group("/users")
		protected_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_user.GET("/my", userController.GetMyProfile)
		}
	}
}
