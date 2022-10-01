package mono

import "errors"

var currencyCodes = map[int32]Currency{
	840: {
		Name:   "US Dollar",
		Code:   "USD",
		Symbol: "$",
	},
	980: {
		Name:   "Hryvnia",
		Code:   "UAH",
		Symbol: "₴",
	},
	978: {
		Name:   "Euro",
		Code:   "EUR",
		Symbol: "€",
	},
	643: {
		Name:   "Russian Ruble",
		Code:   "RUB",
		Symbol: "₽",
	},
	826: {
		Name:   "Pound Sterling",
		Code:   "GBP",
		Symbol: "£",
	},
	756: {
		Name:   "Swiss Franc",
		Code:   "CHF",
		Symbol: "₣",
	},
	933: {
		Name:   "Belarussian Ruble",
		Code:   "BYN",
		Symbol: "Br",
	},
	124: {
		Name:   "Canadian Dollar",
		Code:   "CAD",
		Symbol: "$",
	},
	203: {
		Name:   "Czech Koruna",
		Code:   "CZK",
		Symbol: "Kč",
	},
	208: {
		Name:   "Danish Krone",
		Code:   "DKK",
		Symbol: "Kr",
	},
	348: {
		Name:   "Forint",
		Code:   "HUF",
		Symbol: "Ft",
	},
	985: {
		Name:   "Zloty",
		Code:   "PLN",
		Symbol: "zł",
	},
	949: {
		Name:   "Turkish Lira",
		Code:   "TRY",
		Symbol: "₺",
	},
}

// Currency is internal representation of fiat currencies.
type Currency struct {
	Name   string
	Code   string
	Symbol string
}

// CurrencyFromISO4217 converts ISO4217 to matching currency.
func CurrencyFromISO4217(code int32) (Currency, error) {
	if _, ok := currencyCodes[code]; !ok {
		return Currency{}, errors.New("code is not valid")
	}

	return currencyCodes[code], nil
}

// Exchange contains market buy/sell rates.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
type Exchange struct {
	CodeA     int32   `json:"currencyCodeA"`
	CodeB     int32   `json:"currencyCodeB"`
	Date      int32   `json:"date"`
	RateSell  float64 `json:"rateSell"`
	RateBuy   float64 `json:"rateBuy"`
	RateCross float64 `json:"rateCross"`
}

// Base returns normal representation of CurrencyCodeA.
func (ex *Exchange) Base() (Currency, error) {
	return CurrencyFromISO4217(ex.CodeA)
}

// Quote returns normal representation of CurrencyCodeB.
func (ex *Exchange) Quote() (Currency, error) {
	return CurrencyFromISO4217(ex.CodeB)
}
