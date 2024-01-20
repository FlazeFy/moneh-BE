package httphandlers

import (
	middlewares "moneh/middlewares/jwt"
	"moneh/modules/auth/models"
	"moneh/modules/auth/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func PostLoginUser(c echo.Context) error {
	var body models.UserLogin
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := middlewares.CheckLogin(c, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostRegister(c echo.Context) error {
	var body models.UserRegister
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := repositories.PostUserRegister(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
