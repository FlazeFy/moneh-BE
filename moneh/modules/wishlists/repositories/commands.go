package repositories

import (
	"moneh/modules/wishlists/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func HardDelWishlistById(id, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	validateOwner, err := builders.ValidateOwner(con, baseTable, token, id)
	if err != nil {
		return res, err
	}

	if validateOwner {
		// Command builder
		sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, "id")

		// Exec
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id)
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
	} else {
		// Response
		res.Status = http.StatusNotFound
		res.Message = generator.GenerateQueryMsg(baseTable, 0)
		res.Data = nil
	}

	return res, nil
}

func SoftDelWishlistById(id, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	validateOwner, err := builders.ValidateOwner(con, baseTable, token, id)
	if err != nil {
		return res, err
	}

	if validateOwner {
		// Command builder
		sqlStatement = builders.GetTemplateCommand("soft_delete", baseTable, "id")

		// Exec
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
	} else {
		// Response
		res.Status = http.StatusNotFound
		res.Message = generator.GenerateQueryMsg(baseTable, 0)
		res.Data = nil
	}

	return res, nil
}

func PostWishlist(d models.PostWishlist, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	userId, err := builders.GetUserIdFromToken(con, token)
	if err != nil {
		return res, err
	}

	if userId != "" {
		// Data
		id := uuid.Must(uuid.NewRandom())

		// Command builder
		sqlStatement = "INSERT INTO " + baseTable + " (id, wishlists_name, wishlists_desc, wishlists_img_url, wishlists_type, wishlists_priority, wishlists_price, is_achieved, created_at, created_by, updated_at, deleted_at) " +
			"VALUES (?,?,?,?,?,?,?,?,?,null,null)"

		// Exec
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id, d.WishlistName, d.WishlistDesc, d.WishlistImgUrl, d.WishlistType, d.WishlistPriority, d.WishlistPrice, d.IsAchieved, dt, userId)
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
	} else {
		// Response
		res.Status = http.StatusNotFound
		res.Message = generator.GenerateQueryMsg(baseTable, 0)
		res.Data = nil
	}

	return res, nil
}
