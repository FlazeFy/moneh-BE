package pocket

import (
	"moneh/middlewares"
	"moneh/middlewares/business"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func PocketRouter(r *gin.Engine, pocketController PocketController, redisClient *redis.Client, currencyMiddleware *business.CurrencyMiddleware, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_pocket_all := api.Group("/pockets")
		protected_pocket_all.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_pocket_all.Use(currencyMiddleware.CurrencyMiddleware())
			protected_pocket_all.GET("/", pocketController.GetAllPocket)
		}

		// Private Routes - User
		protected_pocket_user := api.Group("/pockets")
		protected_pocket_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_pocket_user.POST("/", pocketController.CreatePocket, middlewares.AuditTrailMiddleware(db, "post_create_pocket"))
		}
	}
}
