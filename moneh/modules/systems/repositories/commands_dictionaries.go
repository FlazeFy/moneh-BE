package repositories

import (
	"moneh/modules/systems/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
)

func HardDelDictionaryById(id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

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

	return res, nil
}

func PostDictionary(d models.PostDictionaryByType) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, dictionaries_type, dictionaries_name) " +
		"VALUES (null,?,?)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(d.DctType, d.DctName)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	id, err := result.LastInsertId()
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
