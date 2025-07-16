package models

type (
	StatsContextTotal struct {
		Total   int    `json:"total"`
		Context string `json:"context"`
	}
	StatsFlowMonthly struct {
		TotalIncome   int    `json:"total_income"`
		TotalSpending int    `json:"total_spending"`
		Context       string `json:"context"`
	}
)
