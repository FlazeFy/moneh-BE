package repositories

import (
	"database/sql"
	"moneh/modules/users/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/converter"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"strings"
)

func GetMyProfile(token string) (response.Response, error) {
	// Declaration
	var obj models.GetMyProfile
	var res response.Response
	var baseTable = "users"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)

	// Nullable Column
	var LastName sql.NullString
	var ImageUrl sql.NullString
	var AcceptedAt sql.NullString
	var TelegramUserId sql.NullString

	// Query builder
	join := builders.GetTemplateJoin("total", baseTable, "id", baseTable+"_tokens", "context_id", false)

	sqlStatement = "SELECT " + baseTable + ".id, username, first_name, last_name, email, image_url, telegram_user_id, telegram_is_valid, accepted_at " +
		"FROM " + baseTable + " " +
		join + " " +
		"WHERE token = '" + token + "' " +
		"LIMIT 1"

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
			&obj.ID,
			&obj.Username,
			&obj.FirstName,
			&LastName,
			&obj.Email,
			&ImageUrl,
			&TelegramUserId,
			&obj.TelegramIsValid,
			&AcceptedAt,
		)

		if err != nil {
			return res, err
		}
		// Nullable check
		obj.LastName = converter.CheckNullString(LastName)
		obj.ImageUrl = converter.CheckNullString(ImageUrl)
		obj.TelegramUserId = converter.CheckNullString(TelegramUserId)
		obj.AcceptedAt = converter.CheckNullString(AcceptedAt)
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
		res.Data = obj
	}

	return res, nil
}
