package pocket

import (
	"errors"
	"math"
	"moneh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PocketController struct {
	PocketService PocketService
}

func NewPocketController(pocketService PocketService) *PocketController {
	return &PocketController{PocketService: pocketService}
}

// Queries
func (c *PocketController) GetAllPocket(ctx *gin.Context) {
	// Pagination
	pagination := utils.GetPagination(ctx)

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "pocket", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get All Pocket
	pocket, total, err := c.PocketService.GetAllPocket(pagination, *userID)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "pocket", "get", http.StatusNotFound, nil, nil)
		default:
			utils.BuildErrorMessage(ctx, err.Error())
		}
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	metadata := gin.H{
		"total":       total,
		"page":        pagination.Page,
		"limit":       pagination.Limit,
		"total_pages": totalPages,
	}
	utils.BuildResponseMessage(ctx, "success", "pocket", "get", http.StatusOK, pocket, metadata)
}
