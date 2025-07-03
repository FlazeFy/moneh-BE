package flow

import (
	"moneh/middlewares"
	"moneh/middlewares/business"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func FlowRouter(r *gin.Engine, flowController FlowController, redisClient *redis.Client, currencyMiddleware *business.CurrencyMiddleware) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_flow_all := api.Group("/flows")
		protected_flow_all.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_flow_all.DELETE("/:id", flowController.SoftDeleteById)

			protected_flow_all.Use(currencyMiddleware.CurrencyMiddleware())
			protected_flow_all.GET("/", flowController.GetAllFlow)
		}

		// Private Routes - User
		protected_flow_user := api.Group("/flows")
		protected_flow_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_flow_user.POST("/", flowController.CreateFlow)
		}
	}
}
