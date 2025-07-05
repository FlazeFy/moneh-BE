package pocket

import (
	"errors"
	"math"
	"moneh/models"
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

	// Currency
	currency, exist := ctx.Get("currency")
	if !exist {
		utils.BuildResponseMessage(ctx, "failed", "pocket", "currency not available", http.StatusBadRequest, nil, nil)
		return
	}

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

	// Currency Conversion
	for i := range pocket {
		convertedAmount, err := utils.ConvertFromIDR(pocket[i].PocketAmmount, currency.(string))
		if err == nil {
			pocket[i].PocketAmmount = int(convertedAmount)
		}

		if pocket[i].PocketLimit != nil {
			convertedLimit, err := utils.ConvertFromIDR(*pocket[i].PocketLimit, currency.(string))
			if err == nil {
				intConvertedLimit := int(convertedLimit)
				pocket[i].PocketLimit = &intConvertedLimit
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
	utils.BuildResponseMessage(ctx, "success", "pocket", "get", http.StatusOK, pocket, metadata)
}

func (c *PocketController) CreatePocket(ctx *gin.Context) {
	// Models
	var req models.Pocket

	// Validate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		formattedErrors := utils.BuildValidationError(err)
		utils.BuildResponseMessage(ctx, "failed", "pocket", formattedErrors, http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "pocket", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Create Pocket
	pocket, err := c.PocketService.CreatePocket(&req, *userID)
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "pocket", "post", http.StatusCreated, pocket, nil)
}
