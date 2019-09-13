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
	"time"
)

// Authorizer is an interface for different types of MonoBank API authorization.
type Authorizer interface {
	Auth(request *http.Request) error
}

type authCore struct {
	core
	auth Authorizer
}

func newAuthCore(auth Authorizer) *authCore {
	return &authCore{
		auth: auth,
		core: *newCore(),
	}
}

// GetJSON builds the full endpoint path and gets the raw JSON.
func (ac *authCore) GetJSON(ctx context.Context, endpoint string, headers map[string]string) ([]byte, int, error) {
	uri, err := ac.buildURL(endpoint)
	if err != nil {
		return nil, 0, err
	}

	r, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return nil, 0, err
	}

	if err := ac.auth.Auth(r); err != nil {
		return nil, 0, err
	}

	// Set headers.
	for k, v := range headers {
		r.Header.Set(k, v)
	}

	resp, err := ac.Do(r)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

// PostJSON builds the full endpoint path and gets the raw JSON.
func (ac *authCore) PostJSON(
	ctx context.Context,
	endpoint string,
	headers map[string]string,
	payload io.Reader,
) ([]byte, int, error) {
	uri, err := ac.buildURL(endpoint)
	if err != nil {
		return nil, 0, err
	}

	r, err := http.NewRequestWithContext(ctx, "POST", uri, payload)
	if err != nil {
		return nil, 0, err
	}

	if err := ac.auth.Auth(r); err != nil {
		return nil, 0, err
	}

	// Set headers.
	for k, v := range headers {
		r.Header.Set(k, v)
	}

	resp, err := ac.Do(r)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

// User returns user personal information from MonoBank API.
// See https://api.monobank.ua/docs/#operation--personal-client-info-get for details.
func (ac *authCore) User(ctx context.Context, headers map[string]string) (*UserInfo, error) {
	body, status, err := ac.GetJSON(ctx, "/personal/client-info", headers)
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

	var data UserInfo
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// Transactions returns list of transactions from {from} till {to} time.
// See https://api.monobank.ua/docs/#/definitions/StatementItems for details.
func (ac *authCore) Transactions(
	ctx context.Context,
	account string,
	from, to time.Time,
	headers map[string]string,
) (
	[]Transaction,
	error,
) {
	path := fmt.Sprintf("/personal/statement/%s/%d/%d", account, from.Unix(), to.Unix())
	body, status, err := ac.GetJSON(ctx, path, headers)
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
func (ac *authCore) SetWebHook(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	buff, err := json.Marshal(struct{ WebHookUrl string }{url})
	if err != nil {
		return nil, err
	}

	contents, status, err := ac.PostJSON(ctx, "/personal/webhook", headers, bytes.NewReader(buff))
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
