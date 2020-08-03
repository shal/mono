package mono

import (
	"context"
	"time"

	"shal.dev/mono/auth"
	"shal.dev/mono/option"
)

// Personal gives access to personal methods.
type Personal struct {
	client *Client
}

// NewPersonal returns new client of MonoBank Personal API.
func NewPersonal(token string, opts ...option.ClientOption) (*Personal, error) {
	authorizer := auth.NewPersonal(token)

	client, err := NewClient(authorizer, opts...)
	if err != nil {
		return nil, err
	}

	return &Personal{
		client: client,
	}, nil
}

// User returns user personal information from MonoBank API.
// See https://api.monobank.ua/docs/#operation--personal-client-info-get for details.
func (p *Personal) User(ctx context.Context) (*UserInfo, error) {
	return p.client.User(ctx, nil)
}

// Transactions returns list of transactions from {from} till {to} time.
// See https://api.monobank.ua/docs/#/definitions/StatementItems for details.
func (p *Personal) Transactions(ctx context.Context, account string, from, to time.Time) ([]Transaction, error) {
	return p.client.Transactions(ctx, account, from, to, nil)
}

// SetWebHook sets WebHook URL for authorized user.
// See https://api.monobank.ua/docs#operation--personal-webhook-post for details.
func (p *Personal) SetWebHook(ctx context.Context, url string) ([]byte, error) {
	return p.client.SetWebHook(ctx, url, nil)
}

// Rates returns list of currencies rates from MonoBank API.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
func (p *Personal) Rates(ctx context.Context) ([]Exchange, error) {
	return p.client.Rates(ctx)
}
