package pocket

import (
	"moneh/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func PocketRouter(r *gin.Engine, pocketController PocketController, redisClient *redis.Client) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_pocket_all := api.Group("/pockets")
		protected_pocket_all.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_pocket_all.GET("/", pocketController.GetAllPocket)
		}
	}
}
