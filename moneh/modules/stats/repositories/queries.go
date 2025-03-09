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
	"strings"
)

func GetSummaryAppsRepo() (response.Response, error) {
	// Declaration
	var obj models.GetSummaryAppsModel
	var res response.Response

	// Query builder
	count := builders.GetFormulaQuery(nil, "total_item")
	sqlStatement := "SELECT " +
		"(SELECT " + count + " total FROM users) AS total_user, " +
		"(SELECT " + count + " total FROM wishlists) AS total_wishlist, " +
		"(SELECT " + count + " total FROM pockets) AS total_pockets, " +
		"(SELECT " + count + " total FROM flows) AS total_flows"

	// Exec
	con := database.CreateCon()
	row := con.QueryRow(sqlStatement)

	// Map
	err := row.Scan(&obj.TotalUser, &obj.TotalWishlist, &obj.TotalPocket, &obj.TotalFlow)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", 1)
	res.Data = obj

	return res, nil
}

func GetTotalStatsRepo(ord string, view string, table string, typeStats string, extraTotal, token *string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = table
	var mainCol = view
	var sqlStatement string

	if token != nil {
		tokenStr := strings.Replace(*token, "Bearer ", "", -1)
		token = &tokenStr
	}

	// Converted column
	var totalStr string

	// Query builder
	sqlStatement = builders.GetTemplateStats(mainCol, baseTable, typeStats, ord, extraTotal, token)

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	} else {
		defer rows.Close()
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

func GetDashboardRepo(token string) (response.Response, error) {
	// Declaration
	var obj models.GetDashboard
	var arrobj []models.GetDashboard
	var res response.Response
	var sqlStatement string
	var baseTable = "flows"
	token = strings.Replace(token, "Bearer ", "", -1)

	// Converted column
	var LastSpending sql.NullString
	var LastIncome sql.NullString
	var MostHighestIncome sql.NullString
	var lastIncomeVal sql.NullString
	var lastSpendingVal sql.NullString
	var mostExpensiveSpendingVal sql.NullString
	var mostHighestIncomeVal sql.NullString
	var totalItemIncome sql.NullString
	var totalItemSpending sql.NullString
	var myBalance sql.NullString

	// Query builder
	join := builders.GetTemplateJoin("total", baseTable, "created_by", "users_tokens", "context_id", false)
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
		myBalanceSql + " my_balance FROM " + baseTable + " " +
		join + " " +
		"WHERE token = '" + token + "' " +
		" limit 1"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	} else {
		defer rows.Close()
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&LastIncome,
			&LastSpending,
			&obj.MostExpensiveSpending,
			&MostHighestIncome,
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
		lastIncomeValInt, err := converter.ConvertNullStringToInt(lastIncomeVal)
		if err != nil {
			return res, err
		}
		lastSpendingValInt, err := converter.ConvertNullStringToInt(lastSpendingVal)
		if err != nil {
			return res, err
		}
		mostExpensiveSpendingValInt, err := converter.ConvertNullStringToInt(mostExpensiveSpendingVal)
		if err != nil {
			return res, err
		}
		mostHighestIncomeValInt, err := converter.ConvertNullStringToInt(mostHighestIncomeVal)
		if err != nil {
			return res, err
		}
		totalItemIncomeInt, err := converter.ConvertNullStringToInt(totalItemIncome)
		if err != nil {
			return res, err
		}
		totalItemSpendingInt, err := converter.ConvertNullStringToInt(totalItemSpending)
		if err != nil {
			return res, err
		}
		myBalanceInt, err := converter.ConvertNullStringToInt(myBalance)
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
		obj.LastIncome = converter.CheckNullString(LastIncome)
		obj.MostHighestIncome = converter.CheckNullString(MostHighestIncome)

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", 1)
	res.Data = arrobj

	return res, nil
}
