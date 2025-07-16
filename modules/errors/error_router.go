package errors

import (
	"moneh/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ErrorRouter(r *gin.Engine, errorController ErrorController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - Admin
		protected_error_admin := api.Group("/errors")
		protected_error_admin.Use(middlewares.AuthMiddleware(redisClient, "admin"))
		{
			protected_error_admin.GET("/", errorController.GetAllError)
			protected_error_admin.DELETE("/destroy/:id", errorController.HardDeleteErrorById)
		}
	}
}
