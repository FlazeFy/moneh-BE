package flow

import (
	"moneh/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func FlowRouter(r *gin.Engine, flowController FlowController, redisClient *redis.Client) {
	api := r.Group("/api/v1")
	{
		// Private Routes
		protected_flow := api.Group("/flows")
		protected_flow.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_flow.GET("/", flowController.GetAllFlow)
			protected_flow.DELETE("/:id", flowController.SoftDeleteById)
		}
	}
}
