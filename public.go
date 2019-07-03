package mono

import (
	"encoding/json"
	"fmt"
)

// Rates returns list of currencies rates from Monobank API.
// See https://api.monobank.ua/docs/#/definitions/FCurrencyInfo for details.
func (c *Client) Rates() ([]Exchange, error) {
	contents, status, err := c.GetJSON("/bank/currency")
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("invalid status %d", status)
	}

	var data []Exchange
	if err = json.Unmarshal(contents, &data); err != nil {
		return nil, err
	}

	return data, nil
}
