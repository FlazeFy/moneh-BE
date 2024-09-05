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
	result, err := repositories.GetAllFlow(page, 10, "api/v1/flows/"+ord, ord)
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
	result, err := repositories.GetSummaryByType("api/v1/flows/summary/"+types, types)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelFlowById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelFlowById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SoftDelFlowById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.SoftDelFlowById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostFlow(c echo.Context) error {
	var obj models.GetFlow

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

	result, err := repositories.PostFlow(obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalItemAmmountPerDateByType(c echo.Context) error {
	types := c.Param("type")
	view := c.Param("view")
	result, err := repositories.GetTotalItemAmmountPerDateByType(types, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetMonthlyFlowItem(c echo.Context) error {
	mon := c.Param("mon")
	year := c.Param("year")
	types := c.Param("type")

	result, err := repositories.GetMonthlyFlowItem(mon, year, types)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetMonthlyFlowTotal(c echo.Context) error {
	mon := c.Param("mon")
	year := c.Param("year")
	types := c.Param("type")

	result, err := repositories.GetMonthlyFlowTotal(mon, year, types)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
