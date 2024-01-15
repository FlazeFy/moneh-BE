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
	"strconv"
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

	// Converted column
	var WishlistPrice string

	// Query builder
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "wishlists_name", ord)

	sqlStatement = "SELECT id, wishlists_name, wishlists_desc, wishlists_img_url, wishlists_type, wishlists_priority, wishlists_price " +
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
			&WishlistPrice,
		)

		if err != nil {
			return res, err
		}

		// Nullable check
		obj.WishlistImgUrl = converter.CheckNullString(WishlistImgUrl)

		// Converted
		intWishlistPrice, err := strconv.Atoi(WishlistPrice)
		if err != nil {
			return res, err
		}

		obj.WishlistPrice = intWishlistPrice

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

func GetSummary(path string) (response.Response, error) {
	// Declaration
	var obj models.GetSummary
	var arrobj []models.GetSummary
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string

	// Converted Column
	var Average string
	var Achieved string
	var TotalItem string
	var TotalAmmount string

	// Query builder
	col := "wishlists_price"
	col2 := "is_achieved = 1"
	col3 := "wishlists_type"
	avgQuery := builders.GetFormulaQuery(&col, "average")
	totalItemConQuery := builders.GetFormulaQuery(&col2, "total_condition")
	totalItemQuery := builders.GetFormulaQuery(nil, "total_item")
	totalAmmountQuery := builders.GetFormulaQuery(&col, "total_sum")
	exp := builders.GetFormulaQuery(&col, "max")
	chp := builders.GetFormulaQuery(&col, "min")
	mostType := builders.GetFormulaQuery(&col3, "max")

	sqlStatement = "SELECT " + avgQuery + " average, " +
		totalItemConQuery + " achieved, " +
		totalItemQuery + " total_item, " +
		totalAmmountQuery + " total_ammount, " +
		exp + " most_expensive, " +
		chp + " cheapest, " +
		mostType + " most_type " +
		"FROM " + baseTable

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
			&Achieved,
			&TotalItem,
			&TotalAmmount,
			&obj.MostExpensive,
			&obj.Cheapest,
			&obj.MostType,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intAverage, err := strconv.Atoi(Average)
		intAchieved, err := strconv.Atoi(Achieved)
		intTotalItem, err := strconv.Atoi(TotalItem)
		intTotalAmmount, err := strconv.Atoi(TotalAmmount)
		if err != nil {
			return res, err
		}

		obj.Average = intAverage
		obj.Achieved = intAchieved
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
