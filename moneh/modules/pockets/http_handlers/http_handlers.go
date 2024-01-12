package httphandlers

import (
	"moneh/modules/pockets/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllPocketHeaders(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	result, err := repositories.GetAllPocketHeaders(page, 10, "api/v1/pockets/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostPocket(c echo.Context) error {
	result, err := repositories.PostPocket(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelPocketById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelPocketById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
