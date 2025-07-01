package flow

import (
	"errors"
	"math"
	"moneh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FlowController struct {
	FlowService FlowService
}

func NewFlowController(flowService FlowService) *FlowController {
	return &FlowController{FlowService: flowService}
}

// Queries
func (c *FlowController) GetAllFlow(ctx *gin.Context) {
	// Pagination
	pagination := utils.GetPagination(ctx)

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "flow", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get All Flow
	flow, total, err := c.FlowService.GetAllFlow(pagination, *userID)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "flow", "get", http.StatusNotFound, nil, nil)
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
	utils.BuildResponseMessage(ctx, "success", "flow", "get", http.StatusOK, flow, metadata)
}
