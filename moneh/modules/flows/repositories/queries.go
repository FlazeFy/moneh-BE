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
	"strings"
)

func GetAllFlow(page, pageSize int, path string, ord string, token string) (response.Response, error) {
	// Declaration
	var obj models.GetFlow
	var arrobj []models.GetFlow
	var res response.Response
	var baseTable = "flows"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)

	// Converted Column
	var FlowsAmmount string

	// Nullable Column
	var FlowsTag sql.NullString
	var UpdatedAt sql.NullString

	// Query builder
	activeTemplate := builders.GetTemplateLogic("active")
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "flows_name", ord)
	join := builders.GetTemplateJoin("total", baseTable, "created_by", "users_tokens", "context_id", false)

	sqlStatement = "SELECT " + baseTable + ".id, flows_type, flows_category, flows_name, flows_desc, flows_ammount, flows_tag, is_shared, " + baseTable + ".created_at, updated_at " +
		"FROM " + baseTable + " " +
		join + " " +
		"WHERE " + activeTemplate + " " +
		"AND token = '" + token + "' " +
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
			&obj.CreatedAt,
			&UpdatedAt,
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
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)

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

func GetSummaryByType(path, types, token string) (response.Response, error) {
	// Declaration
	var obj models.GetSummaryByType
	var arrobj []models.GetSummaryByType
	var res response.Response
	var baseTable = "flows"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)

	// Converted Column
	var Average string
	var TotalItem string
	var TotalAmmount string

	// Query builder
	col := "flows_ammount"
	avgQuery := builders.GetFormulaQuery(&col, "average")
	totalItemQuery := builders.GetFormulaQuery(nil, "total_item")
	totalAmmountQuery := builders.GetFormulaQuery(&col, "total_sum")
	join := builders.GetTemplateJoin("total", baseTable, "created_by", "users_tokens", "context_id", false)

	sqlStatement = "SELECT " + avgQuery + " average, " +
		totalItemQuery + " total_item, " +
		totalAmmountQuery + " total_ammount " +
		"FROM " + baseTable + " " +
		join + " " +
		"WHERE flows_type = '" + types + "' " +
		"AND token = '" + token + "' "

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

func GetTotalItemAmmountPerDateByType(types, view, token string) (response.Response, error) {
	// Declaration
	var obj models.GetTotalItemAmmountPerDateByType
	var arrobj []models.GetTotalItemAmmountPerDateByType
	var res response.Response
	var baseTable = "flows"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)

	// Converted Column
	var TotalItem string
	var TotalAmmount string

	// Query builder
	col := map[string]string{
		"to_count": "flows_ammount",
		"to_get":   baseTable + ".created_at",
		"view":     view,
	}
	query := converter.MapToString(col)
	queryFinal := builders.GetFormulaQuery(&query, "periodic")
	join := builders.GetTemplateJoin("total", baseTable, "created_by", "users_tokens", "context_id", false)

	sqlStatement = "SELECT " + queryFinal + " " +
		"FROM " + baseTable + " " +
		join + " " +
		"WHERE flows_type = '" + types + "' " +
		"AND token = '" + token + "' " +
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

func GetMonthlyFlowItem(mon, year, types, token string) (response.Response, error) {
	// Declaration
	var obj models.GetMonthlyFlow
	var arrobj []models.GetMonthlyFlow
	var res response.Response
	var baseTable = "flows"
	token = strings.Replace(token, "Bearer ", "", -1)

	var flowWhere = ""
	join := builders.GetTemplateJoin("total", baseTable, "created_by", "users_tokens", "context_id", false)

	if types != "all" {
		flowWhere = "AND flows_type = '" + types + "'"
	}

	sqlStatement := "SELECT flows_name as title, DATE(" + baseTable + ".created_at) as context " +
		"FROM " + baseTable + " " +
		join + " " +
		"WHERE MONTH(" + baseTable + ".created_at) = '" + mon + "' " +
		"AND token = '" + token + "' " +
		"AND YEAR(" + baseTable + ".created_at) = '" + year + "' " +
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

func GetMonthlyFlowTotal(mon, year, types, token string) (response.Response, error) {
	// Declaration
	var obj models.GetMonthlyFlowTotal
	var arrobj []models.GetMonthlyFlowTotal
	var res response.Response
	var baseTable = "flows"
	token = strings.Replace(token, "Bearer ", "", -1)

	var flowWhere = ""
	join := builders.GetTemplateJoin("total", baseTable, "created_by", "users_tokens", "context_id", false)

	if types != "final_total" {
		flowWhere = "AND flows_type = '" + types + "'"
	}

	sqlStatement := "SELECT SUM(flows_ammount) as title, flows_type as type, DATE(" + baseTable + ".created_at) as context " +
		"FROM " + baseTable + " " +
		join + " " +
		"WHERE MONTH(" + baseTable + ".created_at) = '" + mon + "' " +
		"AND token = '" + token + "' " +
		"AND YEAR(" + baseTable + ".created_at) = '" + year + "' " +
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
			&obj.Type,
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

func GetAllFlowExport() (response.Response, error) {
	// Declaration
	var obj models.GetFlowExport
	var arrobj []models.GetFlowExport
	var res response.Response
	var baseTable = "flows"
	var sqlStatement string

	// Converted Column
	var FlowsAmmount string

	// Query builder
	activeTemplate := builders.GetTemplateLogic("active")

	sqlStatement = "SELECT flows_type, flows_category, flows_name, flows_desc, flows_ammount, created_at " +
		"FROM " + baseTable + " " +
		"WHERE " + activeTemplate + " " +
		"ORDER BY created_at DESC "

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
			&obj.FlowsType,
			&obj.FlowsCategory,
			&obj.FlowsName,
			&obj.FlowsDesc,
			&FlowsAmmount,
			&obj.CreatedAt,
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

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	res.Data = arrobj

	return res, nil
}
