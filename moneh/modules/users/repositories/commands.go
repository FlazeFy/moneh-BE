package repositories

import (
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func UpdateTelegram(data echo.Context, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "users"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	token = strings.Replace(token, "Bearer ", "", -1)

	// Query builder
	join := builders.GetTemplateJoin("total", baseTable, "id", baseTable+"_tokens", "context_id", false)

	// Data
	telegramUserId := data.FormValue("telegram_user_id")
	telegramIsValid := data.FormValue("telegram_is_valid")

	// Command builder
	sqlStatement = "UPDATE " + baseTable + " " + join + " SET telegram_user_id=?, telegram_is_valid=?, updated_at=? " +
		"WHERE token = '" + token + "' "

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(telegramUserId, telegramIsValid, dt)
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

	return res, nil
}
