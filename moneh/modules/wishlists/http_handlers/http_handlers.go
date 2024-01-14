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
	result, err := repositories.GetAllWishlistHeaders(page, 10, "api/v1/wishlists/headers/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelWishlistById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelWishlistById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SoftDelWishlistById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.SoftDelWishlistById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostWishlist(c echo.Context) error {
	result, err := repositories.PostWishlist(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
