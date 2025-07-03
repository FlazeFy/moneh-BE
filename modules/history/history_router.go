package history

import (
	"moneh/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func HistoryRouter(r *gin.Engine, historyController HistoryController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_history_all := api.Group("/histories")
		protected_history_all.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_history_all.GET("/my", historyController.GetMyHistory)
		}

		// Private Routes - User
		protected_history_user := api.Group("/histories")
		protected_history_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_history_user.DELETE("/:id", historyController.HardDeleteHistoryById, middlewares.AuditTrailMiddleware(db, "hard_delete_history_by_id"))
		}
	}
}
