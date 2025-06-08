package auth

import (
	"moneh/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService AuthService
}

func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ac *AuthController) Register(c *gin.Context) {
	// Model
	var req models.User

	// Validator
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	// Register Token
	token, err := ac.AuthService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	// Response
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"status":  "success",
		"data": gin.H{
			"access_token": token,
		},
	})
}

func (ac *AuthController) Login(c *gin.Context) {
	// Model
	var req *models.UserAuth

	// Validator
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	// Token Generate
	token, role, err := ac.AuthService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"message": "user login successfully",
		"status":  "success",
		"data": gin.H{
			"role":         role,
			"access_token": token,
		},
	})
}
