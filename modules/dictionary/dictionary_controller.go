package dictionary

import (
	"errors"
	"math"
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
