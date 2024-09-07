package httphandlers

import (
	"moneh/modules/pockets/models"
	"moneh/modules/pockets/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllPocketHeaders(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetAllPocketHeaders(page, 10, "api/v1/pockets/"+ord, ord, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllPocketExport(c echo.Context) error {
	result, err := repositories.GetAllPocketExport()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostPocket(c echo.Context) error {
	var obj models.GetPocketHeaders

	pocketLimitInt, _ := strconv.Atoi(c.FormValue("pockets_limit"))

	obj.PocketsName = c.FormValue("pockets_name")
	obj.PocketsDesc = c.FormValue("pockets_desc")
	obj.PocketsType = c.FormValue("pockets_type")
	obj.PocketsLimit = pocketLimitInt

	result, err := repositories.PostPocket(obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePocket(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.UpdatePocket(c, id)
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
