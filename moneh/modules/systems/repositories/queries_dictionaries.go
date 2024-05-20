package repositories

import (
	"context"
	"log"
	"math"
	"moneh/modules/systems/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"moneh/packages/utils/pagination"
	"net/http"
	"strconv"
)

func GetDictionaryByType(page, pageSize int, path string, dctType string) (response.Response, error) {
	// Declaration
	var obj models.GetDictionaryByType
	var arrobj []models.GetDictionaryByType
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string
	var where string

	// Converted column
	var ID string

	// Query builder
	if dctType != "all" {
		where = "dictionaries_type = '" + dctType + "' "
	} else {
		where = "1 "
	}

	order := "dictionaries_name DESC "

	sqlStatement = "SELECT id, dictionaries_name " +
		"FROM " + baseTable + " " +
		"WHERE " + where +
		"ORDER BY " + order +
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
			&ID,
			&obj.DctName,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intID, err := strconv.Atoi(ID)
		if err != nil {
			return res, err
		}

		obj.ID = intID

		arrobj = append(arrobj, obj)
	}

	if dctType == "all" {
		// Page
		total, err := builders.GetTotalCount(con, baseTable, &where)
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
	} else {
		res.Status = http.StatusOK
		res.Message = generator.GenerateQueryMsg(baseTable, 1)
		res.Data = arrobj
	}

	return res, nil
}

func GetDictionaryByTypeFirebase() (response.Response, error) {
	// Declaration
	var obj map[string]models.GetDictionaryByType
	var res response.Response
	var baseTable = "dictionaries"

	// Exec
	ctx := context.Background()
	client, err := database.InitializeFirebaseDB(ctx)
	if err != nil {
		log.Fatalln("error in initializing firebase DB client: ", err)
		return res, err
	}

	ref := client.NewRef(baseTable)
	if err := ref.Get(ctx, &obj); err != nil {
		log.Fatalln("error in reading from firebase DB:", err)
		return res, err
	}

	var arrobj []models.GetDictionaryByType
	for _, v := range obj {
		arrobj = append(arrobj, v)
	}

	// Build response
	total := len(arrobj)
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
