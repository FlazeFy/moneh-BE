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

	result, err := repositories.GetTotalStats(ord, view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalFlowByCat(c echo.Context) error {
	ord := c.Param("ord")
	view := "flows_category"
	table := "flows"

	result, err := repositories.GetTotalStats(ord, view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalDctByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "dictionaries_type"
	table := "dictionaries"

	result, err := repositories.GetTotalStats(ord, view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalPocketByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "pockets_type"
	table := "pockets"

	result, err := repositories.GetTotalStats(ord, view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalAmmountFlowByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "flows_type"
	table := "flows"
	ext := "flows_ammount"

	result, err := repositories.GetTotalStats(ord, view, table, "total_ammount", &ext)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalWishlistType(c echo.Context) error {
	ord := c.Param("ord")
	view := "wishlists_type"
	table := "wishlists"

	result, err := repositories.GetTotalStats(ord, view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalWishlistPriority(c echo.Context) error {
	ord := c.Param("ord")
	view := "wishlists_priority"
	table := "wishlists"

	result, err := repositories.GetTotalStats(ord, view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalWishlistIsAchieved(c echo.Context) error {
	ord := c.Param("ord")
	view := "is_achieved"
	table := "wishlists"

	result, err := repositories.GetTotalStats(ord, view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalDictionaryToModule(c echo.Context) error {
	table := c.Param("table")
	col := c.Param("col")

	result, err := repositories.GetTotalDictionaryToModule(table, col)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetDashboard(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetDashboard(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
