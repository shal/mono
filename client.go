package mono

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// DefaultBaseURL is production URL of Monobank API.
const DefaultBaseURL = "https://api.monobank.ua"

// BaseURL is a customizable URL of API.
var BaseURL = DefaultBaseURL

// Client is the core structure for Monobank API access.
type Client struct {
	http.Client
	token string
}

func (c *Client) buildURL(endpoint string) (string, error) {
	baseURL, err := url.Parse(BaseURL)
	if err != nil {
		return "", err
	}

	baseURL.Path = path.Join(baseURL.Path, endpoint)
	return baseURL.String(), nil
}

// New creates a new Monobank client with some reasonable HTTP request defaults.
func New(token string) *Client {
	return &Client{
		token: token,
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
func (c *Client) GetJSON(endpoint string) ([]byte, int, error) {
	uri, err := c.buildURL(endpoint)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("X-Token", c.token)

	resp, err := c.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}
