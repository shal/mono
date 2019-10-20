package mono

import (
	"context"
	"net/http"
	"time"
)

type personalAuth struct {
	token string
}

func (auth *personalAuth) Auth(request *http.Request) error {
	request.Header.Set("X-Token", auth.token)
	return nil
}

func newPersonalAuth(token string) Authorizer {
	return &personalAuth{token}
}

// Personal gives access to personal methods.
type Personal struct {
	authCore
}

// NewPersonal returns new client of MonoBank Personal API.
func NewPersonal(token string) *Personal {
	return &Personal{
		authCore: *newAuthCore(newPersonalAuth(token)),
	}
}

func (p *Personal) WithContext(context context.Context) Personal {
	newP := *p
	newP.context = context
	return newP
}

// User returns user personal information from MonoBank API.
// See https://api.monobank.ua/docs/#operation--personal-client-info-get for details.
func (p *Personal) User() (*UserInfo, error) {
	if p.context == nil {
		return p.authCore.User(nil)
	} else {
		return p.authCore.User(nil)
	}
}

// Transactions returns list of transactions from {from} till {to} time.
// See https://api.monobank.ua/docs/#/definitions/StatementItems for details.
func (p *Personal) Transactions(account string, from, to time.Time) ([]Transaction, error) {
	if p.context == nil {
		return p.authCore.Transactions(account, from, to, nil)
	} else {
		return p.authCore.Transactions(account, from, to, nil)
	}
}

// SetWebHook sets WebHook URL for authorized user.
// See https://api.monobank.ua/docs#operation--personal-webhook-post for details.
func (p *Personal) SetWebHook(url string) ([]byte, error) {
	if p.context == nil {
		return p.authCore.SetWebHook(url, nil)
	} else {
		return p.authCore.SetWebHook(url, nil)
	}
}

// Rates returns list of currencies rates from MonoBank API.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
func (p *Personal) Rates() ([]Exchange, error) {
	if p.context == nil {
		return p.authCore.Rates()
	} else {
		return p.authCore.Rates()
	}
}
