package httphandlers

import (
	"moneh/modules/wishlists/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllWishlistHeaders(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	result, err := repositories.GetAllWishlistHeaders(page, 10, "api/v1/wishlists/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
