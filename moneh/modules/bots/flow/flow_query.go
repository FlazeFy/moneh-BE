package flow

import (
	"fmt"
	"moneh/packages/builders"
	"moneh/packages/database"
	"strconv"
	"strings"
)

func GetAllFlowBot() (string, error) {
	// Declaration
	var obj GetAllFlowModel
	var arrobj []GetAllFlowModel
	var baseTable = "flows"
	var sqlStatement string
	var res strings.Builder

	// Converted Column
	var FlowsAmmount string

	// Query builder
	activeTemplate := builders.GetTemplateLogic("active")
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "flows_name", "desc")

	sqlStatement = "SELECT flows_type, flows_category, flows_name, flows_ammount, created_at " +
		"FROM " + baseTable + " " +
		"WHERE " + activeTemplate + " " +
		"ORDER BY " + order

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return "", err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.FlowsType,
			&obj.FlowsCategory,
			&obj.FlowsName,
			&FlowsAmmount,
			&obj.CreatedAt,
		)

		if err != nil {
			return "", err
		}

		// Converted
		intFlowAmmount, err := strconv.Atoi(FlowsAmmount)
		if err != nil {
			return "", err
		}

		obj.FlowsAmmount = intFlowAmmount

		arrobj = append(arrobj, obj)
	}

	for _, flow := range arrobj {
		res.WriteString(fmt.Sprintf(`
				Type : %s
				Category : %s
				Name : %s
				Amount : Rp. %d, 00
				Created At : %s
			`,
			flow.FlowsType,
			flow.FlowsCategory,
			flow.FlowsName,
			flow.FlowsAmmount,
			flow.CreatedAt,
		))
	}

	return res.String(), nil
}
