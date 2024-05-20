package repositories

import (
	"moneh/modules/systems/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/converter"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"

	"github.com/google/uuid"
)

func HardDelTagById(id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "tags"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, "id")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	// Response builder
	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Exec - Firebase
	firebaseDelete := database.DeleteFirebase(id, baseTable)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "permanently delete", int(rowsAffected))
	res.Data = map[string]interface{}{
		"rows_affected":  rowsAffected,
		"is_realtime_db": firebaseDelete,
	}

	return res, nil
}

func PostTag(d models.GetTags) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "tags"
	var sqlStatement string

	// Data
	id := uuid.Must(uuid.NewRandom())

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, tags_slug, tags_name) " +
		"VALUES (?,?,?)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, d.TagSlug, d.TagName)
	if err != nil {
		return res, err
	}

	// Response Builder
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

	return res, nil
}
