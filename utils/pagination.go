package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int
	Limit int
}

func GetPagination(c *gin.Context) (Pagination, error) {
	page, errPage := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, errLimit := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Validation
	if errPage != nil || page < 1 {
		return Pagination{}, fmt.Errorf("page must be an integer greater than 0")
	}
	if errLimit != nil || limit < 1 {
		return Pagination{}, fmt.Errorf("limit must be an integer greater than 0")
	}

	return Pagination{
		Page:  page,
		Limit: limit,
	}, nil
}
