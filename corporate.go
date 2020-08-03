package mono

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"shal.dev/mono/auth"
	"shal.dev/mono/option"
)

const (
	// StatementPermission provides access to balance and statement itself.
	StatementPermission byte = 's'
	// PersonalPermission provides access to client's name and surname.
	PersonalPermission byte = 'p'
)

// Corporate gives access to corporate methods.
type Corporate struct {
	client *Client
}

// NewCorporate returns new client of MonoBank Corporate API.
func NewCorporate(keyData []byte, opts ...option.ClientOption) (*Corporate, error) {
	authorizer, err := auth.NewCorporateAuth(keyData)
	if err != nil {
		return nil, err
	}

	client, err := NewClient(authorizer, opts...)
	if err != nil {
		return nil, err
	}

	return &Corporate{
		client: client,
	}, nil
}

// Auth initializes access.
//func (c *Corporate) Auth(ctx context.Context, callback string, permissions ...byte) (*TokenRequest, error) {
//	timestamp := strconv.Itoa(int(time.Now().Unix()))
//	pp := string(permissions)
//	endpoint := "/personal/auth/request"
//	strings.Join(timestamp, pp, endpoint)
//	sign, err := c.client.auth.Sign(
//	if err != nil {
//		return nil, err
//	}
//
//	headers := map[string]string{
//		"X-Permissions": pp,
//		"X-Sign":        sign,
//		"X-Callback":    callback,
//	}
//
//	body, status, err := c.client.PostJSON(ctx, endpoint, headers, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var tokenRequest TokenRequest
//	if err := json.Unmarshal(body, &tokenRequest); err != nil {
//		return nil, err
//	}
//
//	if status != http.StatusOK {
//		var msg Error
//		if err := json.Unmarshal(body, &msg); err != nil {
//			return nil, errors.New("invalid error payload")
//		}
//		return nil, msg
//	}
//
//	return &tokenRequest, nil
//}

//// CheckAuth checks status of request for client's personal data.
//func (c *Corporate) CheckAuth(ctx context.Context, reqID string) (bool, error) {
//	timestamp := strconv.Itoa(int(time.Now().Unix()))
//	endpoint := "/personal/auth/request"
//
//	sign, err := c.client.auth.Sign(timestamp, reqID, endpoint)
//	if err != nil {
//		return false, err
//	}
//
//	headers := map[string]string{
//		"X-Sign":       sign,
//		"X-Request-Id": reqID,
//	}
//
//	body, status, err := c.client.GetJSON(ctx, endpoint, headers)
//	if err != nil {
//		return false, err
//	}
//
//	if status != http.StatusOK {
//		var msg Error
//		if err := json.Unmarshal(body, &msg); err != nil {
//			return false, errors.New("invalid error payload")
//		}
//		return false, msg
//	}
//
//	return true, nil
//}

// User returns user personal information from MonoBank API.
// See https://api.monobank.ua/docs/#/definitions/UserInfo for details.
func (c *Corporate) User(ctx context.Context, reqID string) (*UserInfo, error) {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	endpoint := "/personal/client-info"

	sign, err := c.client.auth.Sign(timestamp, reqID, endpoint)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"X-Sign":       sign,
		"X-Request-Id": reqID,
	}

	return c.client.User(ctx, headers)
}

// Transactions returns list of transactions from {from} till {to} time.
// See https://api.monobank.ua/docs/#/definitions/StatementItems for details.
func (c *Corporate) Transactions(ctx context.Context, reqID string, account string, from, to time.Time) ([]Transaction, error) {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	fmt.Println()
	path := fmt.Sprintf("/personal/statement/%s/%d/%d", account, from.Unix(), to.Unix())

	sign, err := c.client.auth.Sign(timestamp, reqID, path)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"X-Sign":       sign,
		"X-Request-Id": reqID,
	}

	return c.client.Transactions(ctx, account, from, to, headers)
}

// Rates returns list of currencies rates from MonoBank API.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
func (c *Corporate) Rates(ctx context.Context) ([]Exchange, error) {
	return c.client.Rates(ctx)
}

// GetJSON builds the full endpoint path and gets the raw JSON.
func (c *Corporate) GetJSON(ctx context.Context, endpoint string, headers map[string]string) ([]byte, int, error) {
	return c.client.GetJSON(ctx, endpoint, headers)
}

// PostJSON builds the full endpoint path and gets the raw JSON.
func (c *Corporate) PostJSON(ctx context.Context, endpoint string, headers map[string]string, payload io.Reader) ([]byte, int, error) {
	return c.client.PostJSON(ctx, endpoint, headers, payload)
}
