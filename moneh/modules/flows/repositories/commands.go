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

func HardDelFlowById(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "flows"
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

func SoftDelFlowById(id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "flows"
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

func PostFlow(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "flows"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	flowType := data.FormValue("flows_type")
	flowCat := data.FormValue("flows_category")
	flowName := data.FormValue("flows_name")
	flowDesc := data.FormValue("flows_desc")
	flowAmmount := data.FormValue("flows_ammount")
	flowTag := data.FormValue("flows_tag")
	isShared := data.FormValue("is_shared")

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, flows_type, flows_category, flows_name, flows_desc, flows_ammount, flows_tag, is_shared, created_at, updated_at, deleted_at) " +
		"VALUES (?,?,?,?,?,?,?,?,?,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, flowType, flowCat, flowName, flowDesc, flowAmmount, flowTag, isShared, dt)
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
