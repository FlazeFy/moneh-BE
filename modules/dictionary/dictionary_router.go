package dictionary

import (
	"moneh/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func DictionaryRouter(r *gin.Engine, dictionaryController DictionaryController, redisClient *redis.Client) {
	api := r.Group("/api/v1")
	{
		// Private Routes
		protected_dictionary := api.Group("/dictionaries")
		protected_dictionary.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_dictionary.GET("/", dictionaryController.GetAllDictionary)
		}
	}
}
