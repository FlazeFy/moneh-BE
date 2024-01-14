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

func GetTotalPocketByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "pockets_type"
	table := "pockets"

	result, err := repositories.GetTotalStats("api/v1/stats/pockettype/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalWishlistType(c echo.Context) error {
	ord := c.Param("ord")
	view := "wishlists_type"
	table := "wishlists"

	result, err := repositories.GetTotalStats("api/v1/stats/wishlisttype/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalWishlistPriority(c echo.Context) error {
	ord := c.Param("ord")
	view := "wishlists_priority"
	table := "wishlists"

	result, err := repositories.GetTotalStats("api/v1/stats/wishlistpriority/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalWishlistIsAchieved(c echo.Context) error {
	ord := c.Param("ord")
	view := "is_achieved"
	table := "wishlists"

	result, err := repositories.GetTotalStats("api/v1/stats/wishlistisachieved/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
