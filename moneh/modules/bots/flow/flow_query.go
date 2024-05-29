package flow

import (
	"database/sql"
	"fmt"
	"moneh/modules/stats/models"
	"moneh/packages/builders"
	"moneh/packages/database"
	"moneh/packages/helpers/converter"
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
	var total int
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

		// Calculated
		total += obj.FlowsAmmount

		arrobj = append(arrobj, obj)
	}

	for _, dt := range arrobj {
		amount := converter.ConvertPriceNumber(dt.FlowsAmmount)

		res.WriteString(fmt.Sprintf(`
				Type : %s
				Category : %s
				Name : %s
				Amount : Rp. %s,00
				Created At : %s
			`,
			dt.FlowsType,
			dt.FlowsCategory,
			dt.FlowsName,
			amount,
			dt.CreatedAt,
		))
	}

	// Subtotal
	totalAmount := converter.ConvertPriceNumber(total)
	res.WriteString(fmt.Sprintf(`
			==============================
			Total Amount: Rp. %s,00
		`, totalAmount))

	return res.String(), nil
}

func GetDashboard() (string, error) {
	// Declaration
	var obj models.GetDashboard
	var arrobj []models.GetDashboard
	var sqlStatement string
	var baseTable = "flows"
	var res strings.Builder

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
		return "", err
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
			return "", err
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
			return "", err
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

	for _, dt := range arrobj {
		lastIncomeConverted := converter.ConvertPriceNumber(dt.LastIncomeVal)
		lastSpendingConverted := converter.ConvertPriceNumber(dt.LastSpendingVal)
		mostExpensiveSpendingValConverted := converter.ConvertPriceNumber(dt.MostExpensiveSpendingVal)
		mostHighestIncomeValConverted := converter.ConvertPriceNumber(dt.MostHighestIncomeVal)
		myBalanceConverted := converter.ConvertPriceNumber(dt.MyBalance)

		res.WriteString(fmt.Sprintf(`
			Dashboard

			Last Income 				: %s / Rp. %s,00
			Last Spending 				: %s / Rp. %s,00
			Most Expensive Spending 	: %s / Rp. %s,00
			Most Highest Income 		: %s / Rp. %s,00
			
			Total Item 					: %d Income / %d Spending
			My Balance 					: Rp. %s,00
		`,
			dt.LastIncome,
			lastIncomeConverted,
			dt.LastSpending,
			lastSpendingConverted,
			dt.MostExpensiveSpending,
			mostExpensiveSpendingValConverted,
			dt.MostHighestIncome,
			mostHighestIncomeValConverted,
			dt.TotalItemIncome,
			dt.TotalItemSpending,
			myBalanceConverted,
		))
	}

	return res.String(), nil
}
