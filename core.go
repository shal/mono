package mono

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// DefaultBaseURL is production URL of Monobank API.
const DefaultBaseURL = "https://api.monobank.ua"

type core struct {
	http.Client

	baseURL string
}

func (c *core) buildURL(endpoint string) (string, error) {
	baseURL, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	baseURL.Path = path.Join(baseURL.Path, endpoint)
	return baseURL.String(), nil
}

// newCore creates a new MonoBank client with some reasonable HTTP request defaults.
func newCore() *core {
	return &core{
		baseURL: DefaultBaseURL,
		Client: http.Client{
			Timeout: time.Second * 5,
			Transport: &http.Transport{
				MaxIdleConns:        50,
				MaxIdleConnsPerHost: 50,
			},
		},
	}
}

// GetJSON builds the full endpoint path and gets the raw JSON.
func (c *core) GetJSON(ctx context.Context, endpoint string, headers map[string]string) ([]byte, int, error) {
	uri, err := c.buildURL(endpoint)
	if err != nil {
		return nil, 0, err
	}

	r, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return nil, 0, err
	}

	// Set headers.
	for k, v := range headers {
		r.Header.Set(k, v)
	}

	resp, err := c.Do(r)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

// PostJSON builds the full endpoint path and gets the raw JSON.
func (c *core) PostJSON(
	ctx context.Context,
	endpoint string,
	headers map[string]string,
	payload io.Reader,
) ([]byte, int, error) {
	uri, err := c.buildURL(endpoint)
	if err != nil {
		return nil, 0, err
	}

	r, err := http.NewRequestWithContext(ctx, "POST", uri, payload)
	if err != nil {
		return nil, 0, err
	}

	// Set headers.
	for k, v := range headers {
		r.Header.Set(k, v)
	}

	resp, err := c.Do(r)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

// Rates returns list of currencies rates from MonoBank API.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
func (c *core) Rates(ctx context.Context) ([]Exchange, error) {
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

// SetBaseURL set baseURL to the new specified URL.
func (c *core) SetBaseURL(url string) {
	c.baseURL = url
}
