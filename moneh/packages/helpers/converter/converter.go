package converter

import (
	"database/sql"
	"strings"
)

func CheckNullString(data sql.NullString) string {
	var res string
	if data.Valid {
		res = data.String
	} else {
		res = ""
	}

	return res
}

func TotalChar(val string) int {
	trimed := strings.TrimSpace(val)
	return len(trimed)
}

func ConvertStringBool(val string) bool {
	if val == "0" {
		return false
	} else {
		return true
	}
}
