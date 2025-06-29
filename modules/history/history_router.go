package history

import (
	"github.com/gin-gonic/gin"
)

func HistoryRouter(r *gin.Engine, historyController HistoryController) {
	api := r.Group("/api/v1")
	{
		// Public Routes
		history := api.Group("/histories")
		{
			history.GET("/my", historyController.GetMyHistory)
			history.DELETE("/:id", historyController.HardDeleteHistoryById)
		}
	}
}
