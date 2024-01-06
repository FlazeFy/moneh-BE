package httphandlers

import (
	"moneh/modules/stats/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetTotalFlowByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "flows_type"
	table := "flows"

	result, err := repositories.GetTotalStats("api/v1/stats/flowtype/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalFlowByCat(c echo.Context) error {
	ord := c.Param("ord")
	view := "flows_category"
	table := "flows"

	result, err := repositories.GetTotalStats("api/v1/stats/flowcat/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalDctByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "dictionaries_type"
	table := "dictionaries"

	result, err := repositories.GetTotalStats("api/v1/stats/dcttype/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
