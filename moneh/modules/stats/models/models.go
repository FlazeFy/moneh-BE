package models

type (
	GetMostAppear struct {
		Context string `json:"context"`
		Total   int    `json:"total"`
	}
	GetSummaryAppsModel struct {
		TotalUser     int `json:"total_user"`
		TotalWishlist int `json:"total_wishlist"`
		TotalPocket   int `json:"total_pockets"`
		TotalFlow     int `json:"total_flows"`
	}
	GetDashboard struct {
		LastIncome               string `json:"last_income"`
		LastSpending             string `json:"last_spending"`
		MostExpensiveSpending    string `json:"most_expensive_spending"`
		MostHighestIncome        string `json:"most_highest_income"`
		LastIncomeVal            int    `json:"last_income_value"`
		LastSpendingVal          int    `json:"last_spending_value"`
		MostExpensiveSpendingVal int    `json:"most_expensive_spending_value"`
		MostHighestIncomeVal     int    `json:"most_highest_income_value"`
		TotalItemIncome          int    `json:"total_item_income"`
		TotalItemSpending        int    `json:"total_item_spending"`
		MyBalance                int    `json:"my_balance"`
	}
)
