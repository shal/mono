package mono

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// User returns user personal information from Monobank API.
// See https://api.monobank.ua/docs/#/definitions/UserInfo for details.
func (c *Client) User() (*UserInfo, error) {
	contents, status, err := c.GetJSON("/personal/client-info")
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, fmt.Errorf("invalid status %d", status)
	}

	var data UserInfo
	if err = json.Unmarshal(contents, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// Statement returns list of transactions from {from} till {to} time.
// See https://api.monobank.ua/docs/#/definitions/StatementItems for details.
func (c *Client) Statement(account string, from, to time.Time) ([]StatementItem, error) {
	path := fmt.Sprintf("/personal/statement/%s/%d/%d", account, from.Unix(), to.Unix())
	contents, status, err := c.GetJSON(path)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, fmt.Errorf("invalid status %d", status)
	}

	var data []StatementItem
	if err = json.Unmarshal(contents, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// SetWebHook sets WebHook URL for authorized user.
func (c *Client) SetWebHook(url string) ([]byte, error) {
	buff, err := json.Marshal(struct{ WebHookUrl string }{url})
	if err != nil {
		return nil, err
	}

	contents, status, err := c.PostJSON("/personal/webhook", bytes.NewReader(buff))
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, fmt.Errorf("invalid status %d", status)
	}

	return contents, nil
}
