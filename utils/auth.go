package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserID(ctx *gin.Context) (*uuid.UUID, error) {
	userIDValue, exists := ctx.Get("userID")
	if !exists {
		return nil, errors.New("user not found in context")
	}

	switch v := userIDValue.(type) {
	case string:
		userID, err := uuid.Parse(v)
		if err != nil {
			return nil, err
		}
		return &userID, nil
	case uuid.UUID:
		return &v, nil
	default:
		return nil, errors.New("invalid user id")
	}
}

func GetCurrentRole(c *gin.Context) (string, error) {
	roleVal, exists := c.Get("role")
	if !exists {
		return "", errors.New("role not found in context")
	}

	role, ok := roleVal.(string)
	if !ok {
		return "", errors.New("invalid role format in context")
	}

	return role, nil
}
