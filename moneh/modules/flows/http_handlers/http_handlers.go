package httphandlers

import (
	"moneh/modules/flows/models"
	"moneh/modules/flows/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllFlow(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetAllFlow(page, 10, "api/v1/flows/"+ord, ord, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllFlowExport(c echo.Context) error {
	result, err := repositories.GetAllFlowExport()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetSummaryByType(c echo.Context) error {
	types := c.Param("type")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetSummaryByType("api/v1/flows/summary/"+types, types, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelFlowById(c echo.Context) error {
	id := c.Param("id")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.HardDelFlowById(id, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SoftDelFlowById(c echo.Context) error {
	id := c.Param("id")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.SoftDelFlowById(id, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostFlow(c echo.Context) error {
	var obj models.GetFlow
	token := c.Request().Header.Get("Authorization")

	flowAmmountInt, _ := strconv.Atoi(c.FormValue("flows_ammount"))
	isSharedInt, _ := strconv.Atoi(c.FormValue("is_shared"))

	obj.FlowsType = c.FormValue("flows_type")
	obj.FlowsCategory = c.FormValue("flows_category")
	obj.FlowsName = c.FormValue("flows_name")
	obj.FlowsDesc = c.FormValue("flows_desc")
	obj.FlowsAmmount = flowAmmountInt
	obj.FlowsTag = c.FormValue("flows_tag")
	obj.IsShared = isSharedInt
	obj.CreatedAt = c.FormValue("created_at")

	result, err := repositories.PostFlow(obj, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalItemAmmountPerDateByType(c echo.Context) error {
	types := c.Param("type")
	view := c.Param("view")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetTotalItemAmmountPerDateByType(types, view, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetMonthlyFlowItem(c echo.Context) error {
	mon := c.Param("mon")
	year := c.Param("year")
	types := c.Param("type")
	token := c.Request().Header.Get("Authorization")

	result, err := repositories.GetMonthlyFlowItem(mon, year, types, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetMonthlyFlowTotal(c echo.Context) error {
	mon := c.Param("mon")
	year := c.Param("year")
	types := c.Param("type")
	token := c.Request().Header.Get("Authorization")

	result, err := repositories.GetMonthlyFlowTotal(mon, year, types, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
