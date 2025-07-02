package flow

import (
	"moneh/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func FlowRouter(r *gin.Engine, flowController FlowController, redisClient *redis.Client) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_flow_all := api.Group("/flows")
		protected_flow_all.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_flow_all.GET("/", flowController.GetAllFlow)
			protected_flow_all.DELETE("/:id", flowController.SoftDeleteById)
		}
	}
}
