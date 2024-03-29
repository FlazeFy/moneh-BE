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
	result, err := repositories.GetAllWishlistHeaders(page, 10, "api/v1/wishlists/headers/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetSummary(c echo.Context) error {
	result, err := repositories.GetSummary("api/v1/wishlists/summary")
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
	var obj models.PostWishlist

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

	result, err := repositories.PostWishlist(obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
