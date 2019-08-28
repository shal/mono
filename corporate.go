package mono

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// Statement permission provides access to balance and statement itself.
	StatementPermission byte = 's'
	// Personal permission provides access to client's name and surname.
	PersonalPermission byte = 'p'
)

type CorporateAuth struct {
	*SignTool
	PrivateKey *ecdsa.PrivateKey
	KeyID      string
}

func (auth *CorporateAuth) SignStrings(params ...string) (string, error) {
	return auth.Sign(auth.PrivateKey, strings.Join(params, ""))
}

func (auth *CorporateAuth) Auth(r *http.Request) error {
	timestamp := strconv.Itoa(int(time.Now().Unix()))

	r.Header.Set("X-Key-Id", auth.KeyID)
	r.Header.Set("X-Time", timestamp)

	return nil
}

// Corporate get's access to corporate methods.
type Corporate struct {
	authCore authCore
	auth     CorporateAuth
}

// NewCorporate returns new client of MonoBank Corporate API.
func NewCorporateAuth(
	keyData []byte,
) (*CorporateAuth, error) {
	sign := NewDefaultSignTool()

	privateKey, err := sign.DecodePrivateKey(keyData)
	if err != nil {
		return nil, errors.New("failed to decode private key")
	}

	publicKey := privateKey.PublicKey
	data := elliptic.Marshal(publicKey, publicKey.X, publicKey.Y)
	hash := sha1.New()
	if _, err := hash.Write(data); err != nil {
		return nil, errors.New("failed to encode public key with sha1")
	}
	keyID := hex.EncodeToString(hash.Sum(nil))

	return &CorporateAuth{
		SignTool:   sign,
		PrivateKey: privateKey,
		KeyID:      keyID,
	}, nil
}

func NewCorporate(keyData []byte) (*Corporate, error) {
	auth, err := NewCorporateAuth(keyData)
	if err != nil {
		return nil, err
	}

	return &Corporate{
		auth:     *auth,
		authCore: *newAuthCore(auth),
	}, nil
}

// Auth initializes access.
func (c *Corporate) Auth(callback string, permissions ...byte) (*TokenRequest, error) {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	pp := string(permissions)
	endpoint := "/personal/auth/request"

	sign, err := c.auth.SignStrings(timestamp, pp, endpoint)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"X-Permissions": pp,
		"X-Sign":        sign,
		"X-Callback":    callback,
	}

	body, status, err := c.authCore.PostJSON(endpoint, headers, nil)
	if err != nil {
		return nil, err
	}

	tokenRequest := new(TokenRequest)
	if err := json.Unmarshal(body, tokenRequest); err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		var msg Error
		if err := json.Unmarshal(body, &msg); err != nil {
			return nil, errors.New("invalid error payload")
		}
		return nil, msg
	}

	return tokenRequest, nil
}

// Перевірка статусу запиту на доступ к клієнтським данним
// Check status of auth request.
func (c *Corporate) CheckAuth(reqID string) (bool, error) {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	endpoint := "/personal/auth/request"

	sign, err := c.auth.SignStrings(timestamp, reqID, endpoint)
	if err != nil {
		return false, err
	}

	headers := map[string]string{
		"X-Sign":       sign,
		"X-Request-Id": reqID,
	}

	body, status, err := c.authCore.GetJSON(endpoint, headers)
	if err != nil {
		return false, err
	}

	if status != http.StatusOK {
		var msg Error
		if err := json.Unmarshal(body, &msg); err != nil {
			return false, errors.New("invalid error payload")
		}
		return false, msg
	}

	return true, nil
}

// User returns user personal information from MonoBank API.
// See https://api.monobank.ua/docs/#/definitions/UserInfo for details.
func (c *Corporate) User(reqID string) (*UserInfo, error) {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	endpoint := "/personal/client-info"

	sign, err := c.auth.SignStrings(timestamp, reqID, endpoint)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"X-Sign":       sign,
		"X-Request-Id": reqID,
	}

	return c.authCore.User(headers)
}

// Transactions returns list of transactions from {from} till {to} time.
// See https://api.monobank.ua/docs/#/definitions/StatementItems for details.
func (c *Corporate) Transactions(reqID string, account string, from, to time.Time) ([]Transaction, error) {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	fmt.Println()
	path := fmt.Sprintf("/personal/statement/%s/%d/%d", account, from.Unix(), to.Unix())

	sign, err := c.auth.SignStrings(timestamp, reqID, path)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"X-Sign":       sign,
		"X-Request-Id": reqID,
	}

	return c.authCore.Transactions(account, from, to, headers)
}

// Rates returns list of currencies rates from MonoBank API.
// See https://api.monobank.ua/docs/#/definitions/CurrencyInfo for details.
func (c *Corporate) Rates() ([]Exchange, error) {
	return c.authCore.Rates()
}

// GetJSON builds the full endpoint path and gets the raw JSON.
func (c *Corporate) GetJSON(endpoint string, headers map[string]string) ([]byte, int, error) {
	return c.authCore.GetJSON(endpoint, headers)
}

// PostJSON builds the full endpoint path and gets the raw JSON.
func (c *Corporate) PostJSON(
	endpoint string,
	headers map[string]string,
	payload io.Reader,
) ([]byte, int, error) {
	return c.authCore.PostJSON(endpoint, headers, payload)
}
