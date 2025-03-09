package httphandlers

import (
	"moneh/modules/stats/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetSummaryApps(c echo.Context) error {
	result, err := repositories.GetSummaryAppsRepo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalFlowByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "flows_type"
	table := "flows"
	token := c.Request().Header.Get("Authorization")

	result, err := repositories.GetTotalStatsRepo(ord, view, table, "most_appear", nil, &token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalFlowByCat(c echo.Context) error {
	ord := c.Param("ord")
	view := "flows_category"
	table := "flows"
	token := c.Request().Header.Get("Authorization")

	result, err := repositories.GetTotalStatsRepo(ord, view, table, "most_appear", nil, &token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalPocketByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "pockets_type"
	table := "pockets"
	token := c.Request().Header.Get("Authorization")

	result, err := repositories.GetTotalStatsRepo(ord, view, table, "most_appear", nil, &token)
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
	token := c.Request().Header.Get("Authorization")

	result, err := repositories.GetTotalStatsRepo(ord, view, table, "total_ammount", &ext, &token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalWishlistType(c echo.Context) error {
	ord := c.Param("ord")
	view := "wishlists_type"
	table := "wishlists"
	token := c.Request().Header.Get("Authorization")

	result, err := repositories.GetTotalStatsRepo(ord, view, table, "most_appear", nil, &token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalWishlistPriority(c echo.Context) error {
	ord := c.Param("ord")
	view := "wishlists_priority"
	table := "wishlists"
	token := c.Request().Header.Get("Authorization")

	result, err := repositories.GetTotalStatsRepo(ord, view, table, "most_appear", nil, &token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalWishlistIsAchieved(c echo.Context) error {
	ord := c.Param("ord")
	view := "is_achieved"
	table := "wishlists"
	token := c.Request().Header.Get("Authorization")

	result, err := repositories.GetTotalStatsRepo(ord, view, table, "most_appear", nil, &token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetDashboard(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetDashboardRepo(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
