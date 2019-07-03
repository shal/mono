package mono

type CurrencyInfo struct {
	CurrencyCodeA int32   `json:"CurrencyCodeA"`
	CurrencyCodeB int32   `json:"CurrencyCodeB"`
	Date          int32   `json:"Date"`
	RateSell      float64 `json:"RateSell"`
	RateBuy       float64 `json:"RateBuy"`
	RateCross     float64 `json:"RateCross"`
}
