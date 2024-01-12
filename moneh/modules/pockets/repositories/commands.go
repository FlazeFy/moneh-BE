package repositories

import (
	"moneh/packages/database"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func PostPocket(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "pockets"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	pocketName := data.FormValue("pockets_name")
	pocketDesc := data.FormValue("pockets_desc")
	pocketType := data.FormValue("pockets_type")
	pocketLimit := data.FormValue("pockets_limit")

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, pockets_name, pockets_desc, pockets_type, pockets_limit, created_at, updated_at) " +
		"VALUES (?,?,?,?,?,?,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, pocketName, pocketDesc, pocketType, pocketLimit, dt)
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