package repositories

import (
	"database/sql"
	"math"
	"moneh/modules/pockets/models"
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

func GetAllPocketHeaders(page, pageSize int, path string, ord string, token string) (response.Response, error) {
	// Declaration
	var obj models.GetPocketHeaders
	var arrobj []models.GetPocketHeaders
	var res response.Response
	var baseTable = "pockets"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)

	// Converted Column
	var PocketsLimit string

	// Nullable column
	var UpdatedAt sql.NullString

	// Query builder
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "pockets_name", ord)
	join := builders.GetTemplateJoin("total", baseTable, "created_by", "users_tokens", "context_id", false)

	sqlStatement = "SELECT " + baseTable + ".id, pockets_name, pockets_desc, pockets_type, pockets_limit, " + baseTable + ".created_at, updated_at " +
		"FROM " + baseTable + " " +
		join + " " +
		"WHERE token = '" + token + "' " +
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
			&obj.PocketsName,
			&obj.PocketsDesc,
			&obj.PocketsType,
			&PocketsLimit,
			&obj.CreatedAt,
			&UpdatedAt,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intPocketsLimit, err := strconv.Atoi(PocketsLimit)
		if err != nil {
			return res, err
		}

		obj.PocketsLimit = intPocketsLimit
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

func GetAllPocketExport() (response.Response, error) {
	// Declaration
	var obj models.GetPocketExport
	var arrobj []models.GetPocketExport
	var res response.Response
	var baseTable = "pockets"
	var sqlStatement string

	// Converted Column
	var PocketsLimit string

	sqlStatement = "SELECT pockets_name, pockets_desc, pockets_type, pockets_limit, created_at " +
		"FROM " + baseTable + " " +
		"ORDER BY pockets_limit DESC"

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
			&obj.PocketsName,
			&obj.PocketsDesc,
			&obj.PocketsType,
			&PocketsLimit,
			&obj.CreatedAt,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intPocketsLimit, err := strconv.Atoi(PocketsLimit)
		if err != nil {
			return res, err
		}

		obj.PocketsLimit = intPocketsLimit

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	res.Data = arrobj

	return res, nil
}
