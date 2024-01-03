package repositories

import (
	"moneh/modules/stats/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"strconv"
)

func GetTotalStats(path string, ord string, view string, table string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = table
	var mainCol = view
	var sqlStatement string

	// Converted column
	var totalStr string

	// Query builder
	sqlStatement = builders.GetTemplateStats(mainCol, baseTable, "most_appear", ord, nil)

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
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Total
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
