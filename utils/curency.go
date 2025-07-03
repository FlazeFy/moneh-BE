package utils

import (
	"fmt"
	"moneh/config"
)

func ConvertFromIDR(amount int, targetCurrency string) (float64, error) {
	rate, exists := config.CurrencyRates[targetCurrency]
	if !exists {
		return 0, fmt.Errorf("unsupported currency: %s", targetCurrency)
	}
	converted := float64(amount) * rate

	return converted, nil
}
