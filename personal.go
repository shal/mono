package mono

import (
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

// User returns user personal information from MonoBank API.
// See https://api.monobank.ua/docs/#operation--personal-client-info-get for details.
func (p *Personal) User() (*UserInfo, error) {
	return p.authCore.User(nil)
}

// Transactions returns list of transactions from {from} till {to} time.
// See https://api.monobank.ua/docs/#/definitions/StatementItems for details.
func (p *Personal) Transactions(account string, from, to time.Time) ([]Transaction, error) {
	return p.authCore.Transactions(account, from, to, nil)
}

// SetWebHook sets WebHook URL for authorized user.
// See https://api.monobank.ua/docs#operation--personal-webhook-post for details.
func (p *Personal) SetWebHook(url string) ([]byte, error) {
	return p.authCore.SetWebHook(url, nil)
}

// Rates returns list of currencies rates from MonoBank API.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
func (p *Personal) Rates() ([]Exchange, error) {
	return p.authCore.Rates()
}
