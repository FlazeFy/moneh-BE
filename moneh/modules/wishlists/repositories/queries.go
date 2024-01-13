package repositories

import (
	"database/sql"
	"math"
	"moneh/modules/wishlists/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/converter"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"moneh/packages/utils/pagination"
	"net/http"
)

func GetAllWishlistHeaders(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetWishlistHeaders
	var arrobj []models.GetWishlistHeaders
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string

	// Nullable column
	var WishlistImgUrl sql.NullString

	// Query builder
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "pockets_name")

	sqlStatement = "SELECT id, wishlists_name, wishlists_desc, wishlists_img_url, wishlists_type, wishlists_priority " +
		"FROM " + baseTable + " " +
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
			&obj.WishlistName,
			&obj.WishlistDesc,
			&WishlistImgUrl,
			&obj.WishlistType,
			&obj.WishlistPriority,
		)

		if err != nil {
			return res, err
		}

		// Nullable check
		obj.WishlistImgUrl = converter.CheckNullString(WishlistImgUrl)

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
