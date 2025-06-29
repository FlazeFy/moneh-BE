package feedback

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouteFeedback(r *gin.Engine, db *gorm.DB) {
	// History Module
	feedbackRepo := NewFeedbackRepository(db)
	feedbackService := NewFeedbackService(feedbackRepo)
	feedbackController := NewFeedbackController(feedbackService)

	api := r.Group("/api/v1")
	{
		// Public Routes
		feedback := api.Group("/feedbacks")
		{
			feedback.GET("/", feedbackController.GetAllFeedback)
			feedback.POST("/", feedbackController.CreateFeedback)
		}
	}
}
