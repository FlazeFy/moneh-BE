package feedback

import (
	"github.com/gin-gonic/gin"
)

func FeedbackRouter(r *gin.Engine, feedbackController FeedbackController) {
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
