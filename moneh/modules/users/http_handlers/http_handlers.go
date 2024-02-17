package httphandlers

import (
	"moneh/modules/users/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetMyProfile(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetMyProfile(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
