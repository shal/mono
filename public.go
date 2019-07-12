package mono

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Rates returns list of currencies rates from Monobank API.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
func (c *Client) Rates() ([]Exchange, error) {
	contents, status, err := c.GetJSON("/bank/currency")
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		var msg Error
		if err := json.Unmarshal(contents, &msg); err != nil {
			return nil, errors.New("invalid error payload")
		}
		return nil, msg
	}

	var data []Exchange
	if err = json.Unmarshal(contents, &data); err != nil {
		return nil, err
	}

	return data, nil
}
