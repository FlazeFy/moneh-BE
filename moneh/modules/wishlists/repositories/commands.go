package repositories

import (
	"moneh/modules/wishlists/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func HardDelWishlistById(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, "id")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(slug)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "permanently delete", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func SoftDelWishlistById(id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Command builder
	sqlStatement = builders.GetTemplateCommand("soft_delete", baseTable, "id")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(dt, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "delete", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func PostWishlist(d models.PostWishlist) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, wishlists_name, wishlists_desc, wishlists_img_url, wishlists_type, wishlists_priority, wishlists_price, is_achieved, created_at, updated_at, deleted_at) " +
		"VALUES (?,?,?,?,?,?,?,?,?,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, d.WishlistName, d.WishlistDesc, d.WishlistImgUrl, d.WishlistType, d.WishlistPriority, d.WishlistPrice, d.IsAchieved, dt)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "create", int(rowsAffected))
	res.Data = map[string]interface{}{
		"id":            id,
		"data":          d,
		"rows_affected": rowsAffected,
	}

	return res, nil
}
