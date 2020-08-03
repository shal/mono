package mono

import (
	"shal.dev/mono/iso4217"
)

// Exchange contains market buy/sell rates.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
type Exchange struct {
	CodeA     int32   `json:"CurrencyCodeA"`
	CodeB     int32   `json:"CurrencyCodeB"`
	Date      int32   `json:"Date"`
	RateSell  float64 `json:"RateSell"`
	RateBuy   float64 `json:"RateBuy"`
	RateCross float64 `json:"RateCross"`
}

// Base returns ISO 4217 representation of CurrencyCodeA.
func (ex *Exchange) Base() (*iso4217.Currency, error) {
	return iso4217.CurrencyFromISO4217(ex.CodeA)
}

// Quote returns normal representation of CurrencyCodeB.
func (ex *Exchange) Quote() (*iso4217.Currency, error) {
	return iso4217.CurrencyFromISO4217(ex.CodeB)
}
