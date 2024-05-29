package dct

import (
	"fmt"
	"moneh/packages/database"
)

func GetDctByType(dcttype string) ([]GetDct, error) {
	// Declaration
	var obj GetDct
	var arrobj []GetDct
	var baseTable = "dictionaries"
	var sqlStatement string

	// Query builder
	sqlStatement = "SELECT dictionaries_name " +
		"FROM " + baseTable + " " +
		"WHERE dictionaries_type = '" + dcttype + "' " +
		"ORDER BY dictionaries_name DESC"

	fmt.Println(sqlStatement)

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.DctName,
		)

		if err != nil {
			return nil, err
		}

		arrobj = append(arrobj, obj)
	}

	return arrobj, nil
}
