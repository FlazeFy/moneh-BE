package builders

import (
	"fmt"
	"moneh/packages/helpers/converter"
	"strings"
)

func GetTemplateSelect(name string, firstTable, secondTable *string) string {
	if name == "content_info" {
		return *firstTable + "_slug," + *firstTable + "_name"
	} else if name == "properties_time" {
		return "created_at, created_by"
	} else if name == "properties_full" {
		return "created_at, created_by, updated_at, updated_by"
	} else if name == "user_credential" {
		return "username, email, password, image_url"
	} else if name == "user_mini_info" {
		return "first_name, last_name"
	} else if name == "user_joined_info" {
		return "accepted_at, accepted_by, is_accepted"
	} else if name == "user_access" {
		return "context_type, context_id"
	} else if name == "auth" {
		return "username, password"
	}
	return ""
}

func GetTemplateCommand(name, tableName, colName string) string {
	if name == "soft_delete" {
		return "UPDATE " + tableName + " SET deleted_at = ? WHERE " + tableName + "." + colName + " = ?"
	} else if name == "hard_delete" {
		return "DELETE FROM " + tableName + " WHERE " + tableName + "." + colName + " = ?"
	}
	return ""
}

func GetTemplateConcat(name, col string) string {
	if name == "value_group" {
		return "GROUP_CONCAT(" + col + " SEPARATOR ', ') as " + col
	}
	return ""
}

func GetTemplateOrder(name, tableName, ext, ord string) string {
	if name == "permanent_data" {
		return tableName + ".created_at " + ord + ", " + tableName + "." + ext + " " + ord + " "
	} else if name == "dynamic_data" {
		return tableName + ".updated_at " + ord + ", " + tableName + ".created_at " + ord + ", " + tableName + "." + ext + " " + ord + " "
	} else if name == "most_used_normal" {
		return " COUNT(1) " + ord + ""
	}
	return ""
}

func GetTemplateJoin(typeJoin, firstTable, firstCol, secondTable, secondCol string, nullable bool) string {
	var join = ""
	if nullable {
		join = "LEFT JOIN "
	} else {
		join = "JOIN "
	}
	if typeJoin == "same_col" {
		return join + secondTable + " USING(" + firstCol + ") "
	} else if typeJoin == "total" {
		return join + secondTable + " ON " + secondTable + "." + secondCol + " = " + firstTable + "." + firstCol + " "
	}
	return ""
}

func GetTemplateGroup(is_multi_where bool, col string) string {
	var ext = " WHERE "
	if is_multi_where {
		ext = " AND "
	}

	return ext + col + " IS NOT NULL AND trim(" + col + ") != '' GROUP BY " + col + " "
}

func GetTemplateLogic(name string) string {
	if name == "active" {
		return "deleted_at IS NULL "
	} else if name == "trash" {
		return "deleted_at IS NOT NULL "
	}
	return ""
}

func GetWhereMine(token string) string {
	return "users_tokens.token ='" + token + "'"
}

// Stats
func GetTemplateStats(ctx, firstTable, name string, ord string, joinArgs *string) string {
	// Nullable args
	var args string
	if joinArgs == nil {
		args = ""
	} else {
		args = *joinArgs
	}
	// Notes :
	// Full query
	if name == "most_appear" {
		return "SELECT " + ctx + " as context, " + GetFormulaQuery(nil, "total_item") + " total FROM " + firstTable + " " + args + " GROUP BY " + ctx + " ORDER BY total " + ord
	} else if name == "total_ammount" {
		return "SELECT " + ctx + " as context, " + GetFormulaQuery(joinArgs, "total_sum") + " total FROM " + firstTable + " " + args + " GROUP BY " + ctx + " ORDER BY total " + ord
	}
	return ""
}

func GetFormulaQuery(colTarget *string, name string) string {
	if name == "average" {
		return "CEIL(SUM(" + *colTarget + ") / COUNT(1)) AS "
	} else if name == "total_item" {
		return "COUNT(1) AS "
	} else if name == "total_sum" {
		return "SUM(" + *colTarget + ") AS "
	} else if name == "total_condition" {
		// Column target with condition
		return "COUNT(CASE WHEN " + *colTarget + " THEN 1 END) AS "
	} else if name == "max" {
		return "MAX(" + *colTarget + ") AS "
	} else if name == "min" {
		return "MIN(" + *colTarget + ") AS "
	} else if name == "max_object" || name == "min_object" || name == "total_sum_case" || name == "periodic" {
		prop, err := converter.StringToMap(*colTarget)
		var finalFormulaQuery string
		var whereCount string

		if err != nil {
			fmt.Println("Error:", err)
			return ""
		}

		toCount, _ := prop["to_count"]
		toGet, _ := prop["to_get"]
		fromTable, _ := prop["from_table"]
		whereSyntax, _ := prop["where"]
		joinSyntax, _ := prop["join"]

		finalCount := strings.Split(toCount, " ")
		remainCount := strings.Split(toCount, " AND ")

		var count string = finalCount[0]
		whereCount = " WHERE " + remainCount[0]

		if len(finalCount) > 1 {
			count = finalCount[len(finalCount)-1]
		}

		if name == "max_object" {
			finalFormulaQuery = "(SELECT " + toGet + " FROM " + fromTable
			if joinSyntax != "" {
				finalFormulaQuery += " " + joinSyntax
			}
			finalFormulaQuery += " WHERE " + toCount + " = (SELECT MAX(" + count + ") FROM " + fromTable
			if joinSyntax != "" {
				finalFormulaQuery += " " + joinSyntax
			}
			if whereSyntax != "" {
				finalFormulaQuery += " " + whereSyntax + " AND " + whereCount
			} else {
				finalFormulaQuery += " " + whereCount
			}
			finalFormulaQuery += ") limit 1) AS "
		} else if name == "min_object" {
			finalFormulaQuery = "(SELECT " + toGet + " FROM " + fromTable
			if joinSyntax != "" {
				finalFormulaQuery += " " + joinSyntax
			}
			finalFormulaQuery += " WHERE " + toCount + " = (SELECT MIN(" + count + ") FROM " + fromTable
			if joinSyntax != "" {
				finalFormulaQuery += " " + joinSyntax
			}
			if whereSyntax != "" {
				finalFormulaQuery += " " + whereSyntax + " AND " + whereCount
			} else {
				finalFormulaQuery += " " + whereCount
			}
			finalFormulaQuery += ") limit 1) AS "
		} else if name == "total_sum_case" {
			finalFormulaQuery = "SUM(CASE WHEN " + toGet + " THEN " + toCount + " ELSE 0 END) "
		} else if name == "periodic" {
			finalFormulaQuery = "COUNT(1) as total_item, " + prop["view"] + "(" + toGet + ") as context, SUM(" + toCount + ") as total_ammount"
		}

		return finalFormulaQuery
	}
	return ""
}
