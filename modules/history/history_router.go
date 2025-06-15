package history

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB) {
	// History Module
	historyRepo := NewHistoryRepository(db)
	historyService := NewHistoryService(historyRepo)
	historyController := NewHistoryController(historyService)

	api := r.Group("/api/v1")
	{
		// Public Routes
		history := api.Group("/history")
		{
			history.GET("/all", historyController.GetAllHistory)
			history.GET("/my", historyController.GetMyHistory)
			history.DELETE("/:id", historyController.DeleteHistoryById)
		}
	}
}
