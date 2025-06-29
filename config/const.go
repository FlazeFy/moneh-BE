package config

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
var Currencies = []string{"IDR", "USD", "EUR", "JPY", "GBP", "CNY", "CAD", "CHF", "AUD", "HKD", "SGD"}
