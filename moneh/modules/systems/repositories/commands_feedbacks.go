package repositories

import (
	"moneh/modules/systems/models"
	"moneh/packages/database"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func PostFeedback(d models.PostFeedback) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "feedbacks"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, feedbacks_rate, feedbacks_desc, created_at) " +
		"VALUES (?,?,?,?)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, d.FdbRate, d.FdbDesc, dt)
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
