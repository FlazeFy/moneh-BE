package repositories

import (
	"database/sql"
	"moneh/modules/stats/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/converter"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"net/http"
	"strconv"
)

func GetTotalStats(ord string, view string, table string, typeStats string, extraTotal *string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = table
	var mainCol = view
	var sqlStatement string

	// Converted column
	var totalStr string

	// Query builder
	sqlStatement = builders.GetTemplateStats(mainCol, baseTable, typeStats, ord, extraTotal)

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", 1)
	res.Data = arrobj

	return res, nil
}

func GetTotalDictionaryToModule(table string, col string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable string = "dictionaries"
	var sqlStatement string
	var extraStatement string = "JOIN " + table + " ON dictionaries.dictionaries_name = " + table + "." + col + " " +
		"WHERE dictionaries_type = '" + col + "'"

	// Converted column
	var totalStr string

	// Query builder
	sqlStatement = builders.GetTemplateStats(col, baseTable, "most_appear", "desc", &extraStatement)

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", 1)
	res.Data = arrobj

	return res, nil
}

func GetDashboard() (response.Response, error) {
	// Declaration
	var obj models.GetDashboard
	var arrobj []models.GetDashboard
	var res response.Response
	var sqlStatement string
	var baseTable = "flows"

	// Converted column
	var lastIncomeVal string
	var lastSpendingVal string
	var mostExpensiveSpendingVal string
	var mostHighestIncomeVal string
	var totalItemIncome string
	var totalItemSpending string
	var myBalance string

	// Nullable column
	var LastSpending sql.NullString

	// Query builder
	lastIncomeQueryRaw := map[string]string{
		"to_count":   "flows_type = 'income' AND created_at",
		"to_get":     "flows_name",
		"from_table": "flows",
	}
	lastSpendingQueryRaw := map[string]string{
		"to_count":   "flows_type = 'spending' AND created_at",
		"to_get":     "flows_name",
		"from_table": "flows",
	}
	mostExpensiveSpendingQueryRaw := map[string]string{
		"to_count":   "flows_type = 'spending' AND flows_ammount",
		"to_get":     "flows_name",
		"from_table": "flows",
	}
	mostHighestIncomeQueryRaw := map[string]string{
		"to_count":   "flows_type = 'income' AND flows_ammount",
		"to_get":     "flows_name",
		"from_table": "flows",
	}
	lastIncomeValQueryRaw := map[string]string{
		"to_count":   "flows_type = 'income' AND created_at",
		"to_get":     "flows_ammount",
		"from_table": "flows",
	}
	lastSpendingValQueryRaw := map[string]string{
		"to_count":   "flows_type = 'spending' AND created_at",
		"to_get":     "flows_ammount",
		"from_table": "flows",
	}
	mostExpensiveSpendingValQueryRaw := map[string]string{
		"to_count":   "flows_type = 'spending' AND flows_ammount",
		"to_get":     "flows_ammount",
		"from_table": "flows",
	}
	mostHighestIncomeValQueryRaw := map[string]string{
		"to_count":   "flows_type = 'income' AND flows_ammount",
		"to_get":     "flows_ammount",
		"from_table": "flows",
	}
	myBalanceIncomeQueryRaw := map[string]string{
		"to_count":   "flows_ammount",
		"to_get":     "flows_type = 'income'",
		"from_table": "flows",
	}
	myBalanceSpendingQueryRaw := map[string]string{
		"to_count":   "flows_ammount",
		"to_get":     "flows_type = 'spending'",
		"from_table": "flows",
	}
	tItemConIncome := "flows_type = 'income'"
	tItemConSpend := "flows_type = 'spending'"

	lastIncomeQuery := converter.MapToString(lastIncomeQueryRaw)
	lastSpendingQuery := converter.MapToString(lastSpendingQueryRaw)
	mostExpensiveSpendingQuery := converter.MapToString(mostExpensiveSpendingQueryRaw)
	mostHighestIncomeQuery := converter.MapToString(mostHighestIncomeQueryRaw)
	lastIncomeValQuery := converter.MapToString(lastIncomeValQueryRaw)
	lastSpendingValQuery := converter.MapToString(lastSpendingValQueryRaw)
	mostExpensiveSpendingValQuery := converter.MapToString(mostExpensiveSpendingValQueryRaw)
	mostHighestIncomeValQuery := converter.MapToString(mostHighestIncomeValQueryRaw)
	myBalanceIncomeQuery := converter.MapToString(myBalanceIncomeQueryRaw)
	myBalanceSpendingQuery := converter.MapToString(myBalanceSpendingQueryRaw)

	lastIncomeSql := builders.GetFormulaQuery(&lastIncomeQuery, "max_object")
	lastSpendingSql := builders.GetFormulaQuery(&lastSpendingQuery, "max_object")
	mostExpensiveSpendingSql := builders.GetFormulaQuery(&mostExpensiveSpendingQuery, "max_object")
	mostHighestIncomeSql := builders.GetFormulaQuery(&mostHighestIncomeQuery, "max_object")
	lastIncomeValSql := builders.GetFormulaQuery(&lastIncomeValQuery, "max_object")
	lastSpendingValSql := builders.GetFormulaQuery(&lastSpendingValQuery, "max_object")
	mostExpensiveSpendingValSql := builders.GetFormulaQuery(&mostExpensiveSpendingValQuery, "max_object")
	mostHighestIncomeValSql := builders.GetFormulaQuery(&mostHighestIncomeValQuery, "max_object")
	totalItemIncomeSql := builders.GetFormulaQuery(&tItemConIncome, "total_condition")
	totalItemSpendingSql := builders.GetFormulaQuery(&tItemConSpend, "total_condition")
	myBalanceSql := builders.GetFormulaQuery(&myBalanceIncomeQuery, "total_sum_case") + " - " + builders.GetFormulaQuery(&myBalanceSpendingQuery, "total_sum_case")

	sqlStatement = "SELECT " +
		lastIncomeSql + " last_income, " +
		lastSpendingSql + " last_spending, " +
		mostExpensiveSpendingSql + " most_expensive_spending, " +
		mostHighestIncomeSql + " most_highest_income, " +
		lastIncomeValSql + " last_income_value, " +
		lastSpendingValSql + " last_spending_value, " +
		mostExpensiveSpendingValSql + " most_expensive_spending_value, " +
		mostHighestIncomeValSql + " most_highest_income_value, " +
		totalItemIncomeSql + " total_item_income, " +
		totalItemSpendingSql + " total_item_spending, " +
		myBalanceSql + " my_balance FROM " + baseTable + " limit 1"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.LastIncome,
			&LastSpending,
			&obj.MostExpensiveSpending,
			&obj.MostHighestIncome,
			&lastIncomeVal,
			&lastSpendingVal,
			&mostExpensiveSpendingVal,
			&mostHighestIncomeVal,
			&totalItemIncome,
			&totalItemSpending,
			&myBalance,
		)

		if err != nil {
			return res, err
		}

		// Converted
		lastIncomeValInt, err := strconv.Atoi(lastIncomeVal)
		lastSpendingValInt, err := strconv.Atoi(lastSpendingVal)
		mostExpensiveSpendingValInt, err := strconv.Atoi(mostExpensiveSpendingVal)
		mostHighestIncomeValInt, err := strconv.Atoi(mostHighestIncomeVal)
		totalItemIncomeInt, err := strconv.Atoi(totalItemIncome)
		totalItemSpendingInt, err := strconv.Atoi(totalItemSpending)
		myBalanceInt, err := strconv.Atoi(myBalance)
		if err != nil {
			return res, err
		}

		obj.LastIncomeVal = lastIncomeValInt
		obj.LastSpendingVal = lastSpendingValInt
		obj.MostExpensiveSpendingVal = mostExpensiveSpendingValInt
		obj.MostHighestIncomeVal = mostHighestIncomeValInt
		obj.TotalItemIncome = totalItemIncomeInt
		obj.TotalItemSpending = totalItemSpendingInt
		obj.MyBalance = myBalanceInt

		// Nullable Converter
		obj.LastSpending = converter.CheckNullString(LastSpending)

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", 1)
	res.Data = arrobj

	return res, nil
}
