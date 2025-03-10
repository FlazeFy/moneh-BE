package repositories

import (
	"moneh/modules/flows/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/converter"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func HardDelFlowById(id, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "flows"
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	validateOwner, err := builders.ValidateOwner(con, baseTable, token, id)
	if err != nil {
		return res, err
	}

	if validateOwner {
		// Command builder
		sqlStatement := builders.GetTemplateCommand("hard_delete", baseTable, "id")

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

func SoftDelFlowById(id, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "flows"
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

func PostFlow(d models.GetFlow, tokenorid string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "flows"
	var sqlStatement string
	var dt string
	var marks string
	con := database.CreateCon()

	marks = strings.Replace(tokenorid, "Bearer ", "", -1)
	userId, err := builders.GetUserIdFromToken(con, marks)

	if userId == "" && err == nil {
		marks = userId
	} else if userId != "" && err == nil {
		marks = userId
	} else {
		return res, err
	}

	if d.CreatedAt != "" {
		dt = d.CreatedAt
	} else {
		dt = time.Now().Format("2006-01-02 15:04:05")
	}

	// Data
	id := uuid.Must(uuid.NewRandom())

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, flows_type, flows_category, flows_name, flows_desc, flows_ammount, flows_tag, is_shared, created_at, created_by, updated_at, deleted_at) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?,null,null)"

	// Exec
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, d.FlowsType, d.FlowsCategory, d.FlowsName, d.FlowsDesc, d.FlowsAmmount, d.FlowsTag, d.IsShared, dt, marks)
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

	return res, nil
}
