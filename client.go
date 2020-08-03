package mono

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"shal.dev/mono/auth"
	"shal.dev/mono/option"
)

const (
	MaxIdleConn = 50
	TimeOut     = 30 * time.Second
)

const (
	scheme = "https"
	host   = "api.monobank.ua"
)

type Client struct {
	hc *http.Client

	scheme string
	host   string
	auth   auth.Authorizer
}

func (c *Client) url(endpoint string) string {
	return c.scheme + "://" + c.host + endpoint
}

func NewClient(auth auth.Authorizer, opts ...option.ClientOption) (*Client, error) {
	client := Client{
		scheme: scheme,
		host:   host,
		auth:   auth,
	}

	var settings option.DialSettings
	for _, opt := range opts {
		opt.Apply(&settings)
	}

	// Prepare http.Client.
	if settings.HTTPClient != nil {
		client.hc = settings.HTTPClient
	} else {
		client.hc = &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        MaxIdleConn,
				MaxIdleConnsPerHost: MaxIdleConn,
			},
			Timeout: TimeOut,
		}
	}

	if settings.Endpoint != nil && strings.Contains(*settings.Endpoint, "://") {
		uri, err := url.Parse(*settings.Endpoint)
		if err != nil {
			return nil, fmt.Errorf("endpoint: %w", err)
		}

		client.scheme, client.host = uri.Scheme, uri.Host
	} else if settings.Endpoint != nil {
		client.scheme, client.host = scheme, *settings.Endpoint
	}

	return &client, nil
}

// GetJSON builds the full endpoint path and gets the raw JSON.
func (c *Client) GetJSON(ctx context.Context, endpoint string, headers map[string]string) ([]byte, int, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url(endpoint), nil)
	if err != nil {
		return nil, 0, err
	}

	// Set headers.
	for k, v := range headers {
		r.Header.Set(k, v)
	}

	return c.Do(r)
}

// PostJSON builds the full endpoint path and gets the raw JSON.
func (c *Client) PostJSON(ctx context.Context, endpoint string, headers map[string]string, payload io.Reader) ([]byte, int, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url(endpoint), payload)
	if err != nil {
		return nil, 0, err
	}

	// Set headers.
	for k, v := range headers {
		r.Header.Set(k, v)
	}

	return c.Do(r)
}

// TODO: Add comment.
func (c *Client) Do(r *http.Request) ([]byte, int, error) {
	if c.auth.Sign2(r); err != nil {
		return nil, 0, err
	}

	if c.auth.Auth(r); err != nil {
		return nil, 0, err
	}

	resp, err := c.hc.Do(r)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

// Rates returns list of currencies rates from MonoBank API.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
func (c *Client) Rates(ctx context.Context) ([]Exchange, error) {
	contents, status, err := c.GetJSON(ctx, "/bank/currency", nil)
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

// User returns user personal information from MonoBank API.
// See https://api.monobank.ua/docs/#operation--personal-client-info-get for details.
func (c *Client) User(ctx context.Context, headers map[string]string) (*UserInfo, error) {
	body, status, err := c.GetJSON(ctx, "/personal/client-info", headers)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		var msg Error
		if err := json.Unmarshal(body, &msg); err != nil {
			return nil, fmt.Errorf("unexpected error payload: %w", err)
		}
		return nil, msg
	}

	var data UserInfo
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// Transactions returns list of transactions from {from} till {to} time.
// See https://api.monobank.ua/docs/#/definitions/StatementItems for details.
func (c *Client) Transactions(ctx context.Context, account string, from, to time.Time, headers map[string]string) (
	[]Transaction,
	error,
) {
	path := fmt.Sprintf("/personal/statement/%s/%d/%d", account, from.Unix(), to.Unix())
	body, status, err := c.GetJSON(ctx, path, headers)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		var msg Error
		if err := json.Unmarshal(body, &msg); err != nil {
			return nil, errors.New("invalid error payload")
		}
		return nil, msg
	}

	var data []Transaction
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// SetWebHook sets WebHook URL for authorized user.
// See https://api.monobank.ua/docs#operation--personal-webhook-post for details.
func (c *Client) SetWebHook(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	buff, err := json.Marshal(struct{ WebHookUrl string }{url})
	if err != nil {
		return nil, err
	}

	contents, status, err := c.PostJSON(ctx, "/personal/webhook", headers, bytes.NewReader(buff))
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

	return contents, nil
}


// Auth initializes access.
// TODO: Check that authorizer = corporate.
func (c *Client) Authorize(ctx context.Context, callback string, permissions ...byte) (*TokenRequest, error) {
	if _, ok := c.auth.(*auth.CorporateAuth); !ok {
		return nil, ErrUnexpectedAuth
	}

	pp := string(permissions)

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url("/personal/auth/request"), nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("X-Permissions", pp)
	r.Header.Set("X-Callback", callback)

	err = c.auth.Sign(r)
	if err != nil {
		return nil, err
	}

	err = c.auth.Auth(r)
	if err != nil {
		return nil, err
	}

	body, status, err := c.Do(r)
	if err != nil {
		return nil, err
	}

	var tokenRequest TokenRequest
	if err := json.Unmarshal(body, &tokenRequest); err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		return &tokenRequest, nil
	}

	var msg Error
	if err := json.Unmarshal(body, &msg); err != nil {
		return nil, fmt.Errorf("unexpected error payload: %w", err)
	}

	return nil, msg
}

// CheckAuth checks status of request for client's personal data.
func (c *Client) CheckAuth(ctx context.Context, reqID string) (bool, error) {
	if _, ok := c.auth.(*auth.CorporateAuth); !ok {
		return false, ErrUnexpectedAuth
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url("/personal/auth/request"), nil)
	if err != nil {
		return false, err
	}

	r.Header.Set("X-Request-Id", reqID)

	err = c.auth.Sign(r)
	if err != nil {
		return false, err
	}

	err = c.auth.Auth(r)
	if err != nil {
		return false, err
	}

	body, status, err := c.Do(r)
	if err != nil {
		return false, err
	}

	if status == http.StatusOK {
		return true, nil
	}

	var msg Error
	if err := json.Unmarshal(body, &msg); err != nil {
		return false, errors.New("invalid error payload")
	}

	return false, msg
}