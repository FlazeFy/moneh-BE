package repositories

import (
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
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

func PostWishlist(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	wishlistName := data.FormValue("wishlists_name")
	wishlistDesc := data.FormValue("wishlists_desc")
	wishlistImgUrl := data.FormValue("wishlists_img_url")
	wishlistType := data.FormValue("wishlists_type")
	wishlistPriority := data.FormValue("wishlists_priority")
	wishlistPrice := data.FormValue("wishlists_price")
	isAchieved := data.FormValue("is_achieved")

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, wishlists_name, wishlists_desc, wishlists_img_url, wishlists_type, wishlists_priority, wishlists_price, is_achieved, created_at, updated_at, deleted_at) " +
		"VALUES (?,?,?,?,?,?,?,?,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, wishlistName, wishlistDesc, wishlistImgUrl, wishlistType, wishlistPriority, wishlistPrice, isAchieved, dt)
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
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
