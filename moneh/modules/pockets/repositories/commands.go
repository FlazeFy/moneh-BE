package repositories

import (
	"moneh/modules/pockets/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/converter"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func PostPocket(d models.GetPocketHeaders, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "pockets"
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
		sqlStatement = "INSERT INTO " + baseTable + " (id, pockets_name, pockets_desc, pockets_type, pockets_limit, created_at, created_by, updated_at) " +
			"VALUES (?,?,?,?,?,?,?,null)"

		// Exec
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id, d.PocketsName, d.PocketsDesc, d.PocketsType, d.PocketsLimit, dt, userId)
		if err != nil {
			return res, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return res, err
		}

		// Exec - Firebase
		dataMap, err := converter.StructToMap(d)
		firebaseInsert := database.InsertFirebase(id.String(), baseTable, dataMap)

		// Response
		res.Status = http.StatusOK
		res.Message = generator.GenerateCommandMsg(baseTable, "create", int(rowsAffected))
		res.Data = map[string]interface{}{
			"id":             id,
			"data":           d,
			"rows_affected":  rowsAffected,
			"is_realtime_db": firebaseInsert,
		}
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
		res.Data = nil
	}

	return res, nil
}

func UpdatePocket(data echo.Context, id, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "pockets"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	validateOwner, err := builders.ValidateOwner(con, baseTable, token, id)
	if err != nil {
		return res, err
	}

	if validateOwner {
		// Data
		pocketName := data.FormValue("pockets_name")
		pocketDesc := data.FormValue("pockets_desc")
		pocketType := data.FormValue("pockets_type")
		pocketLimit := data.FormValue("pockets_limit")

		// Command builder
		sqlStatement = "UPDATE " + baseTable + " SET pockets_name=?, pockets_desc=?, pockets_type=?, pockets_limit=?, updated_at=? " +
			"WHERE id=? "

		// Exec
		con := database.CreateCon()
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(pocketName, pocketDesc, pocketType, pocketLimit, dt, id)
		if err != nil {
			return res, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return res, err
		}

		// Response
		res.Status = http.StatusOK
		res.Message = generator.GenerateCommandMsg(baseTable, "update", int(rowsAffected))
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

func HardDelPocketById(id, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "pockets"
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
		con := database.CreateCon()
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
