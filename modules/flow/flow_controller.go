package flow

import (
	"errors"
	"math"
	"moneh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (c *FlowController) SoftDeleteById(ctx *gin.Context) {
	// Param
	id := ctx.Param("id")

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "flow", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Parse Param UUID
	flowID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "flow", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	err = c.FlowService.SoftDeleteFlowById(*userID, flowID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "flow", "empty", http.StatusNotFound, nil, nil)
		default:
			utils.BuildErrorMessage(ctx, err.Error())
		}
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "flow", "soft delete", http.StatusOK, nil, nil)
}
