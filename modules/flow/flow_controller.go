package flow

import (
	"errors"
	"fmt"
	"log"
	"math"
	"moneh/config"
	"moneh/models"
	"moneh/modules/stats"
	"moneh/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FlowController struct {
	FlowService  FlowService
	StatsService stats.StatsService
}

func NewFlowController(flowService FlowService, statsService stats.StatsService) *FlowController {
	return &FlowController{
		FlowService:  flowService,
		StatsService: statsService,
	}
}

// Queries
func (c *FlowController) GetAllFlow(ctx *gin.Context) {
	// Pagination
	pagination := utils.GetPagination(ctx)

	// Currency
	currency, exist := ctx.Get("currency")
	if !exist {
		utils.BuildResponseMessage(ctx, "failed", "flow", "currency not available", http.StatusBadRequest, nil, nil)
		return
	}

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

	// Currency Conversion
	for i := range flow {
		for j := range flow[i].FlowRelations {
			convertedAmount, err := utils.ConvertFromIDR(flow[i].FlowRelations[j].Ammount, currency.(string))
			if err == nil {
				flow[i].FlowRelations[j].Ammount = int(convertedAmount)
			}
		}
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

func (c *FlowController) CreateFlow(ctx *gin.Context) {
	// Models
	var req models.Flow

	// Validate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		formattedErrors := utils.BuildValidationError(err)
		utils.BuildResponseMessage(ctx, "failed", "flow", formattedErrors, http.StatusBadRequest, nil, nil)
		return
	}

	// Validate Flow Relations
	if len(req.FlowRelations) == 0 {
		utils.BuildResponseMessage(ctx, "failed", "flow", "flow_relations must contain at least one item", http.StatusBadRequest, nil, nil)
		return
	}
	for i, rel := range req.FlowRelations {
		if rel.Ammount == 0 {
			utils.BuildResponseMessage(ctx, "failed", "flow", fmt.Sprintf("flow_relation_ammount at index %d is required and must be greater than 0", i), http.StatusBadRequest, nil, nil)
			return
		}
		if rel.PocketId == uuid.Nil {
			utils.BuildResponseMessage(ctx, "failed", "flow", fmt.Sprintf("pocket_id at index %d is required", i), http.StatusBadRequest, nil, nil)
			return
		}
	}

	// Get User ID
	log.Println(req.FlowRelations)
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "flow", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Create Flow
	flow, err := c.FlowService.CreateFlow(&req, *userID)
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "flow", "post", http.StatusCreated, flow, nil)
}

func (c *FlowController) SoftDeleteFlowById(ctx *gin.Context) {
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

func (c *FlowController) HardDeleteFlowById(ctx *gin.Context) {
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

	err = c.FlowService.HardDeleteFlowById(*userID, flowID)
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
	utils.BuildResponseMessage(ctx, "success", "flow", "hard delete", http.StatusOK, nil, nil)
}

func (c *FlowController) GetMostContextFlow(ctx *gin.Context) {
	// Param
	targetCol := ctx.Param("target_col")

	// Validator : Target Column Validator
	if !utils.Contains(config.StatsFlowField, targetCol) {
		utils.BuildResponseMessage(ctx, "failed", "flow", "target_col is not valid", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "flow", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service: Get Most Context
	flow, err := c.StatsService.GetMostUsedContext("flows", targetCol, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "flow", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "flow", "get", http.StatusOK, flow, nil)
}

func (c *FlowController) GetMonthlyFlow(ctx *gin.Context) {
	// Param
	yearStr := ctx.Param("year")

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "flow", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Validator Year
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "flow", "invalid year", http.StatusBadRequest, nil, nil)
		return
	}

	// Service: Get Most Context
	flow, err := c.StatsService.GetMonthlyFlow(year, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "flow", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "flow", "get", http.StatusOK, flow, nil)
}
