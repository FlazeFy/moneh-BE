package business

import (
	"fmt"
	"moneh/modules/user"
	"moneh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Currency Struct
type CurrencyMiddleware struct {
	UserService user.UserService
}

// Currency Constructor
func NewCurrencyMiddleware(userService user.UserService) *CurrencyMiddleware {
	return &CurrencyMiddleware{
		UserService: userService,
	}
}

func (m *CurrencyMiddleware) CurrencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := utils.GetUserID(c)
		if err != nil {
			utils.BuildResponseMessage(c, "failed", "auth", "unauthorized", http.StatusUnauthorized, nil, nil)
			return
		}

		user, err := m.UserService.GetMyProfile(*userID)
		if err != nil {
			utils.BuildResponseMessage(c, "failed", "auth", "user not found", http.StatusNotFound, nil, nil)
			return
		}
		fmt.Println(user.Currency)

		c.Set("currency", user.Currency)
		c.Next()
	}
}
