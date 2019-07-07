package mono

import (
	"net/http"
	"strconv"
	"time"
)

type Authorizer interface {
	Auth(request *http.Request)
}

type PersonalAuthorizer struct {
	token string
}

func (auth *PersonalAuthorizer) Auth(request *http.Request) {
	request.Header.Set("X-Token", auth.token)
}

func NewPersonalAuth(token string) Authorizer {
	return &PersonalAuthorizer{token}
}

type CorporateAuthorizer struct {
	KeyID     string
	RequestID string
	Sign      string
}

func (auth *CorporateAuthorizer) Auth(request *http.Request) {
	request.Header.Set("X-Key-Id", auth.KeyID)
	request.Header.Set("X-Time", strconv.Itoa(int(time.Now().Unix())))
	request.Header.Set("X-Request-Id:", auth.RequestID)
	request.Header.Set("X-Sign", auth.Sign)
}

func NewCorporateAuth(key, id, sign string) Authorizer {
	return &CorporateAuthorizer{key, id, sign}
}

type PublicAuthorizer struct{}

func (auth *PublicAuthorizer) Auth(request *http.Request) {}
