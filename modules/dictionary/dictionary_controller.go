package dictionary

import (
	"errors"
	"math"
	"moneh/models"
	"moneh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DictionaryController struct {
	DictionaryService DictionaryService
}

func NewDictionaryController(dictionaryService DictionaryService) *DictionaryController {
	return &DictionaryController{DictionaryService: dictionaryService}
}

// Queries
func (c *DictionaryController) GetAllDictionary(ctx *gin.Context) {
	// Pagination
	pagination := utils.GetPagination(ctx)

	// Service : Get All Dictionary
	dictionary, total, err := c.DictionaryService.GetAllDictionary(pagination)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "dictionary", "get", http.StatusNotFound, nil, nil)
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
	utils.BuildResponseMessage(ctx, "success", "dictionary", "get", http.StatusOK, dictionary, metadata)
}

// Command
func (c *DictionaryController) CreateDictionary(ctx *gin.Context) {
	// Models
	var req models.Dictionary

	// Validate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BuildResponseMessage(ctx, "failed", "question", "invalid request body", http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Create Dictionary
	dictionary := models.Dictionary{
		DictionaryType: req.DictionaryType,
		DictionaryName: req.DictionaryName,
	}
	err := c.DictionaryService.CreateDictionary(&dictionary)
	if err != nil {
		if err.Error() == "dictionary already exists" {
			utils.BuildResponseMessage(ctx, "failed", "dictionary", "already exists", http.StatusConflict, nil, nil)
			return
		}

		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "dictionary", "post", http.StatusCreated, nil, nil)
}
