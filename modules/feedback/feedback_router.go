package feedback

import (
	"moneh/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func FeedbackRouter(r *gin.Engine, feedbackController FeedbackController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Public Routes
		feedback := api.Group("/feedbacks")
		{
			feedback.POST("/", feedbackController.CreateFeedback)
		}

		// Private Routes - Admin
		protected_feedback_admin := api.Group("/feedbacks")
		protected_feedback_admin.Use(middlewares.AuthMiddleware(redisClient, "admin"))
		{
			protected_feedback_admin.GET("/", feedbackController.GetAllFeedback)
			protected_feedback_admin.DELETE("/destroy/:id", feedbackController.HardDeleteFeedbackById, middlewares.AuditTrailMiddleware(db, "hard_delete_feedback_by_id"))
		}
	}
}
