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
)

func GetAllTags(page, pageSize int, path string) (response.Response, error) {
	// Declaration
	var obj models.GetTags
	var arrobj []models.GetTags
	var res response.Response
	var baseTable = "tags"
	var sqlStatement string

	sqlStatement = "SELECT tags_slug, tags_name " +
		"FROM " + baseTable + " " +
		"ORDER BY tags_name DESC " +
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
			&obj.TagSlug,
			&obj.TagName,
		)

		if err != nil {
			return res, err
		}

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

func GetAllTagsFirebase() (response.Response, error) {
	// Declaration
	var obj map[string]models.GetTags
	var res response.Response
	var baseTable = "tags"

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

	var arrobj []models.GetTags
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
