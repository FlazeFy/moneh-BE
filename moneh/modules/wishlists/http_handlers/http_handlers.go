package httphandlers

import (
	"moneh/modules/wishlists/models"
	"moneh/modules/wishlists/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllWishlistHeaders(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetAllWishlistHeaders(page, 10, "api/v1/wishlists/headers/"+ord, ord, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetSummary(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetSummary("api/v1/wishlists/summary", token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelWishlistById(c echo.Context) error {
	id := c.Param("id")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.HardDelWishlistById(id, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SoftDelWishlistById(c echo.Context) error {
	id := c.Param("id")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.SoftDelWishlistById(id, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostWishlist(c echo.Context) error {
	var obj models.PostWishlist
	token := c.Request().Header.Get("Authorization")

	// Converted
	WishlistPriceInt, _ := strconv.Atoi(c.FormValue("wishlists_price"))
	IsAchievedInt, _ := strconv.Atoi(c.FormValue("is_achieved"))

	obj.WishlistName = c.FormValue("wishlists_name")
	obj.WishlistDesc = c.FormValue("wishlists_desc")
	obj.WishlistImgUrl = c.FormValue("wishlists_img_url")
	obj.WishlistType = c.FormValue("wishlists_type")
	obj.WishlistPriority = c.FormValue("wishlists_priority")
	obj.WishlistPrice = WishlistPriceInt
	obj.IsAchieved = IsAchievedInt

	result, err := repositories.PostWishlist(obj, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
