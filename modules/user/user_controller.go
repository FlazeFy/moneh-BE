package user

import (
	"errors"
	"moneh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserService UserService
}

func NewUserController(historyService UserService) *UserController {
	return &UserController{UserService: historyService}
}

// Queries
func (c *UserController) GetMyProfile(ctx *gin.Context) {
	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "profile", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get My Profile
	profile, err := c.UserService.GetMyProfile(*userID)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "profile", "get", http.StatusNotFound, nil, nil)
		default:
			utils.BuildErrorMessage(ctx, err.Error())
		}
		return
	}

	utils.BuildResponseMessage(ctx, "success", "profile", "get", http.StatusOK, profile, nil)
}
