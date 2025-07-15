package config

import "time"

var ResponseMessages = map[string]string{
	"post":        "created",
	"put":         "updated",
	"hard delete": "permanentally deleted",
	"soft delete": "deleted",
	"recover":     "recovered",
	"get":         "fetched",
	"login":       "login",
	"sign out":    "signed out",
	"empty":       "not found",
}
var CurrencyRates = map[string]float64{
	"IDR": 1.0,
	"USD": 0.000065,
	"EUR": 0.000060,
	"JPY": 0.0095,
	"GBP": 0.000051,
	"CNY": 0.00047,
	"CAD": 0.000088,
	"CHF": 0.000057,
	"AUD": 0.000096,
	"HKD": 0.00051,
	"SGD": 0.000087,
}

// Rules
var RedisTime = 10 * time.Minute

var StatsFlowField = []string{"flow_type", "flow_category"}
var StatsPocketField = []string{"pocket_type"}

var Currencies = []string{"IDR", "USD", "EUR", "JPY", "GBP", "CNY", "CAD", "CHF", "AUD", "HKD", "SGD"}
var FlowTypes = []string{"Income", "Spending"}
var FlowCategories = []string{"Food & Drink", "Transportation", "Entertainment", "Health", "Shopping", "Bills & Utilities", "Education", "Investment", "Salary", "Gift & Donation", "Travel", "Rent", "Insurance", "Pet Care", "Others"}
var PocketTypes = []string{"Bank", "Ovo", "Link Aja", "Go Pay", "Shopee Pay", "Cash", "PayPal", "Dana"}
