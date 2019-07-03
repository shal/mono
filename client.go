package mono

// BaseURL is URL of Monobank API.
var BaseURL = "https://api.monobank.ua"

// Client is the core structure for Monobank API access.
type Client struct {
	token string
}

// TODO: Add description.
func New(token string) *Client {
	return &Client{token}
}

// TODO: Add description.
func (c *Client) Rates() []CurrencyInfo {
	return []CurrencyInfo{}
}
