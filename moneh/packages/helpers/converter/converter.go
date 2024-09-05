package converter

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strconv"
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
func ConvertNullStringToInt(ns sql.NullString) (int, error) {
	if ns.Valid {
		return strconv.Atoi(ns.String)
	}
	return 0, nil
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

func MapToString(val map[string]string) string {
	result, _ := json.Marshal(val)
	return string(result)
}

func StringToMap(val string) (map[string]string, error) {
	var result map[string]string
	err := json.Unmarshal([]byte(val), &result)
	return result, err
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	out := make(map[string]interface{})
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i).Interface()
		fieldName := typ.Field(i).Tag.Get("json")

		// Custom conversion logic
		switch fieldValue.(type) {
		case string:
			// Example of converting string to int
			if intValue, err := strconv.Atoi(fieldValue.(string)); err == nil {
				out[fieldName] = intValue
			} else {
				out[fieldName] = fieldValue
			}
		default:
			out[fieldName] = fieldValue
		}
	}

	return out, nil
}

func ConvertPriceNumber(n int) string {
	in := strconv.Itoa(n)
	out := ""

	for i, r := range in {
		if i > 0 && (len(in)-i)%3 == 0 {
			out += "."
		}
		out += string(r)
	}

	return out
}
