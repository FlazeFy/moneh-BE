package pocket

import (
	"fmt"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/converter"
	"strconv"
	"strings"
)

func GetAllPocketBot() (string, error) {
	// Declaration
	var obj GetAllPocketModel
	var arrobj []GetAllPocketModel
	var baseTable = "pockets"
	var sqlStatement string
	var res strings.Builder

	// Converted Column
	var PocketLimit string

	// Query builder
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "pockets_name", "desc")

	sqlStatement = "SELECT pockets_name, pockets_desc, pockets_type, pockets_limit " +
		"FROM " + baseTable + " " +
		"ORDER BY " + order

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return "", err
	}

	// Map
	var total int
	for rows.Next() {
		err = rows.Scan(
			&obj.PocketsName,
			&obj.PocketsDesc,
			&obj.PocketsType,
			&PocketLimit,
		)

		if err != nil {
			return "", err
		}

		// Converted
		intPocketsLimit, err := strconv.Atoi(PocketLimit)
		if err != nil {
			return "", err
		}

		obj.PocketsLimit = intPocketsLimit

		// Calculated
		total += obj.PocketsLimit

		arrobj = append(arrobj, obj)
	}

	for _, dt := range arrobj {
		limit := converter.ConvertPriceNumber(dt.PocketsLimit)

		res.WriteString(fmt.Sprintf(`
				Type : %s
				Name : %s
				Notes : %s
				Amount : Rp. %s,00
			`,
			dt.PocketsType,
			dt.PocketsName,
			dt.PocketsDesc,
			limit,
		))
	}

	// Subtotal
	totalLimit := converter.ConvertPriceNumber(total)
	res.WriteString(fmt.Sprintf(`
			==============================
			Total Limit: Rp. %s,00
		`, totalLimit))

	return res.String(), nil
}
