package mono

import (
	"encoding/json"
	"errors"
)

type CurrencyInfo struct {
	CurrencyCodeA int32   `json:"CurrencyCodeA"`
	CurrencyCodeB int32   `json:"CurrencyCodeB"`
	Date          int32   `json:"Date"`
	RateSell      float64 `json:"RateSell"`
	RateBuy       float64 `json:"RateBuy"`
	RateCross     float64 `json:"RateCross"`
}

// Rates returns list of currencies rates from Monobank API.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
func (c *Client) Rates() ([]CurrencyInfo, error) {
  contents, status, err := c.GetJSON("/bank/currency")
  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, errors.New("")
  }

  var data []CurrencyInfo
  if err = json.Unmarshal(contents, &data); err != nil {
    return nil, err
  }

  return data, nil
}