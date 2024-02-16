package repositories

import (
	"database/sql"
	"math"
	"moneh/modules/flows/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/converter"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"moneh/packages/utils/pagination"
	"net/http"
	"strconv"
)

func GetAllFlow(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetFlow
	var arrobj []models.GetFlow
	var res response.Response
	var baseTable = "flows"
	var sqlStatement string

	// Converted Column
	var FlowsAmmount string

	// Nullable Column
	var FlowsTag sql.NullString

	// Query builder
	activeTemplate := builders.GetTemplateLogic("active")
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "flows_name", ord)

	sqlStatement = "SELECT id, flows_type, flows_category, flows_name, flows_desc, flows_ammount, flows_tag, is_shared " +
		"FROM " + baseTable + " " +
		"WHERE " + activeTemplate + " " +
		"ORDER BY " + order + " " +
		"LIMIT ? OFFSET ?"

	// Exec
	con := database.CreateCon()
	offset := (page - 1) * pageSize
	rows, err := con.Query(sqlStatement, pageSize, offset)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Id,
			&obj.FlowsType,
			&obj.FlowsCategory,
			&obj.FlowsName,
			&obj.FlowsDesc,
			&FlowsAmmount,
			&FlowsTag,
			&obj.IsShared,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intFlowAmmount, err := strconv.Atoi(FlowsAmmount)
		if err != nil {
			return res, err
		}

		obj.FlowsAmmount = intFlowAmmount
		obj.FlowsTag = converter.CheckNullString(FlowsTag)

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, path)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = map[string]interface{}{
			"current_page":   page,
			"data":           arrobj,
			"first_page_url": pagination.FirstPageURL,
			"from":           pagination.From,
			"last_page":      pagination.LastPage,
			"last_page_url":  pagination.LastPageURL,
			"links":          pagination.Links,
			"next_page_url":  pagination.NextPageURL,
			"path":           pagination.Path,
			"per_page":       pageSize,
			"prev_page_url":  pagination.PrevPageURL,
			"to":             pagination.To,
			"total":          total,
		}
	}

	return res, nil
}

func GetSummaryByType(path string, types string) (response.Response, error) {
	// Declaration
	var obj models.GetSummaryByType
	var arrobj []models.GetSummaryByType
	var res response.Response
	var baseTable = "flows"
	var sqlStatement string

	// Converted Column
	var Average string
	var TotalItem string
	var TotalAmmount string

	// Query builder
	col := "flows_ammount"
	avgQuery := builders.GetFormulaQuery(&col, "average")
	totalItemQuery := builders.GetFormulaQuery(nil, "total_item")
	totalAmmountQuery := builders.GetFormulaQuery(&col, "total_sum")

	sqlStatement = "SELECT " + avgQuery + " average, " +
		totalItemQuery + " total_item, " +
		totalAmmountQuery + " total_ammount " +
		"FROM " + baseTable + " " +
		"WHERE flows_type = '" + types + "' "

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&Average,
			&TotalItem,
			&TotalAmmount,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intAverage, err := strconv.Atoi(Average)
		intTotalItem, err := strconv.Atoi(TotalItem)
		intTotalAmmount, err := strconv.Atoi(TotalAmmount)
		if err != nil {
			return res, err
		}

		obj.Average = intAverage
		obj.TotalItem = intTotalItem
		obj.TotalAmmount = intTotalAmmount

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetTotalItemAmmountPerDateByType(types, view string) (response.Response, error) {
	// Declaration
	var obj models.GetTotalItemAmmountPerDateByType
	var arrobj []models.GetTotalItemAmmountPerDateByType
	var res response.Response
	var baseTable = "flows"
	var sqlStatement string

	// Converted Column
	var TotalItem string
	var TotalAmmount string

	// Query builder
	col := map[string]string{
		"to_count": "flows_ammount",
		"to_get":   "created_at",
		"view":     view,
	}
	query := converter.MapToString(col)
	queryFinal := builders.GetFormulaQuery(&query, "periodic")

	sqlStatement = "SELECT " + queryFinal + " " +
		"FROM " + baseTable + " " +
		"WHERE flows_type = '" + types + "' " +
		"GROUP BY 2 " +
		"LIMIT 7"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&TotalItem,
			&obj.Context,
			&TotalAmmount,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intTotalItem, err := strconv.Atoi(TotalItem)
		intTotalAmmount, err := strconv.Atoi(TotalAmmount)
		if err != nil {
			return res, err
		}

		obj.TotalItem = intTotalItem
		obj.TotalAmmount = intTotalAmmount

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetMonthlyFlowItem(mon, year, types string) (response.Response, error) {
	// Declaration
	var obj models.GetMonthlyFlow
	var arrobj []models.GetMonthlyFlow
	var res response.Response
	var baseTable = "flows"

	var flowWhere = ""

	if types != "all" {
		flowWhere = "AND flows_type = '" + types + "'"
	}

	sqlStatement := "SELECT flows_name as title, DATE(created_at) as context " +
		"FROM " + baseTable + " " +
		"WHERE MONTH(created_at) = '" + mon + "' " +
		"AND YEAR(created_at) = '" + year + "' " +
		flowWhere + " "

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Title,
			&obj.Context,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	if len(arrobj) == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetMonthlyFlowTotal(mon, year, types string) (response.Response, error) {
	// Declaration
	var obj models.GetMonthlyFlow
	var arrobj []models.GetMonthlyFlow
	var res response.Response
	var baseTable = "flows"

	var flowWhere = ""

	if types != "all" {
		flowWhere = "AND flows_type = '" + types + "'"
	}

	sqlStatement := "SELECT SUM(flows_ammount) as title, DATE(created_at) as context " +
		"FROM " + baseTable + " " +
		"WHERE MONTH(created_at) = '" + mon + "' " +
		"AND YEAR(created_at) = '" + year + "' " +
		flowWhere + " " +
		"GROUP BY 2"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Title,
			&obj.Context,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	if len(arrobj) == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
