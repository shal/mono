package iso4217

import (
	"errors"
)

var currencyCodes = map[int32]*Currency{
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
	360: {
		Name:   "Rupiah",
		Code:   "IDR",
		Symbol: "Rp",
	},
	376: {
		Name:   "New Israeli Sheqel",
		Code:   "ILS",
		Symbol: "",
	},
	356: {
		Name:   "Indian Rupee",
		Code:   "INR",
		Symbol: "",
	},
	368: {
		Name:   "Iraqi Dinar",
		Code:   "",
		Symbol: "",
	},
	364: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	352: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	400: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	404: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	417: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	116: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	408: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	410: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	414: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	398: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	418: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	422: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	144: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	434: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	504: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	498: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	969: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	807: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	496: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	478: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	480: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	454: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	484: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	458: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	943: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	516: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	566: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	558: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	578: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	524: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	554: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	512: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	604: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	608: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	586: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	600: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	634: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	946: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	941: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	682: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	690: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	938: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	752: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	702: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	694: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	706: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	968: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	760: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	748: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	764: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	972: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	795: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	788: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	156: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	784: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	971: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	8: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	51: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	973: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	32: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	36: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	944: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	50: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	975: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	48: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	108: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	96: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	68: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	986: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
	72: {
		Name:   "",
		Code:   "",
		Symbol: "",
	},
}

// Currency is internal representation of fiat currencies.
type Currency struct {
	Name   string
	Code   string
	Symbol string
}

// CurrencyFromISO4217 converts ISO4217 to matching currency.
func CurrencyFromISO4217(code int32) (*Currency, error) {
	if _, ok := currencyCodes[code]; !ok {
		return nil, errors.New("code is not valid")
	}

	return currencyCodes[code], nil
}
