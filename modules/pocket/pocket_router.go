package pocket

import (
	"moneh/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func PocketRouter(r *gin.Engine, pocketController PocketController, redisClient *redis.Client) {
	api := r.Group("/api/v1")
	{
		// Private Routes
		protected_pocket := api.Group("/pockets")
		protected_pocket.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_pocket.GET("/", pocketController.GetAllPocket)
		}
	}
}
