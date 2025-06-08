package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCurrentUserID(c *gin.Context) (uuid.UUID, error) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		return uuid.UUID{}, errors.New("user id not found in context")
	}

	userID, err := uuid.Parse(userIDVal.(string))
	if err != nil {
		return uuid.UUID{}, errors.New("user id is not a valid UUID")
	}

	return userID, nil
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
