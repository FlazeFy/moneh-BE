package history

import (
	"math"
	"moneh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HistoryController struct {
	HistoryService HistoryService
}

func NewHistoryController(historyService HistoryService) *HistoryController {
	return &HistoryController{HistoryService: historyService}
}

// @Summary      Get All History
// @Description  Returns a paginated list of all users histories
// @Tags         History
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.ResponseGetAllHistory
// @Failure      404  {object}  map[string]string
// @Router       /api/v1/history/all [get]
func (rc *HistoryController) GetAllHistory(c *gin.Context) {
	// Pagination
	pagination := utils.GetPagination(c)

	// Service: Get All History
	history, total, err := rc.HistoryService.GetAllHistory(pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	// Response
	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	c.JSON(http.StatusOK, gin.H{
		"message": "history fetched",
		"status":  "success",
		"data":    history,
		"metadata": gin.H{
			"total":       total,
			"page":        pagination.Page,
			"limit":       pagination.Limit,
			"total_pages": totalPages,
		},
	})
}

func (rc *HistoryController) GetMyHistory(c *gin.Context) {
	// Pagination
	pagination := utils.GetPagination(c)

	// Get User Id
	userId, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	// Get Role
	role, err := utils.GetCurrentRole(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	// Service: Get My History
	history, total, err := rc.HistoryService.GetMyHistory(pagination, *userId, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	// Response
	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	c.JSON(http.StatusOK, gin.H{
		"message": "history fetched",
		"status":  "success",
		"data":    history,
		"metadata": gin.H{
			"total":       total,
			"page":        pagination.Page,
			"limit":       pagination.Limit,
			"total_pages": totalPages,
		},
	})
}

func (rc *HistoryController) DeleteHistoryById(c *gin.Context) {
	// Get Role
	role, err := utils.GetCurrentRole(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "role not permission",
			"status":  "failed",
		})
		return
	}

	// Param
	id := c.Param("id")

	// Parse Id
	historyID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid UUID format",
			"status":  "failed",
		})
		return
	}

	// Service: Delete History By Id
	err = rc.HistoryService.DeleteHistoryById(historyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"message": "history deleted",
		"status":  "success",
	})
}
