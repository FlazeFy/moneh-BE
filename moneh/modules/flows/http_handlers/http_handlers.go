package httphandlers

import (
	"moneh/modules/flows/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllFlow(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	result, err := repositories.GetAllFlow(page, 10, "api/v1/flows/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelFlowById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelFlowById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SoftDelFlowById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.SoftDelFlowById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostFlow(c echo.Context) error {
	result, err := repositories.PostFlow(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
