package feedback

import (
	"errors"
	"math"
	"moneh/models"
	"moneh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeedbackController struct {
	FeedbackService FeedbackService
}

func NewFeedbackController(feedbackService FeedbackService) *FeedbackController {
	return &FeedbackController{FeedbackService: feedbackService}
}

// Queries
func (c *FeedbackController) GetAllFeedback(ctx *gin.Context) {
	// Pagination
	pagination, err := utils.GetPagination(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Service : Get All Feedback
	feedbackList, total, err := c.FeedbackService.GetAllFeedback(pagination)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "feedback", "empty", http.StatusNotFound, nil, nil)
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
	utils.BuildResponseMessage(ctx, "success", "feedback", "get", http.StatusOK, feedbackList, metadata)
}

// Command
func (c *FeedbackController) CreateFeedback(ctx *gin.Context) {
	// Models
	var req models.Feedback

	// Validate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		formattedErrors := utils.BuildValidationError(err)
		utils.BuildResponseMessage(ctx, "failed", "feedback", formattedErrors, http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "feedback", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Add Feedback
	err = c.FeedbackService.CreateFeedback(&req, *userID)
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "feedback", "post", http.StatusCreated, nil, nil)
}

func (c *FeedbackController) HardDeleteFeedbackById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Parse Param UUID
	feedbackID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "feedback", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Hard Delete Feedback By ID
	err = c.FeedbackService.HardDeleteFeedbackByID(feedbackID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "feedback", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "feedback", "hard delete", http.StatusOK, nil, nil)
}
