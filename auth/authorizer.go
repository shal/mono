package auth

import (
	"net/http"
)

// Authorizer is an interface for different types of MonoBank API authorization.
type Authorizer interface {
	Auth(r *http.Request) error
	Sign(r *http.Request) error
}

type NopAuthorizer struct{}

func (a *NopAuthorizer) Auth(*http.Request) error {
	return nil
}

func (a *NopAuthorizer) Sign(...string) (string, error) {
	return "", nil
}

// NopAuth returns function, that skips authorization.
func NopAuth() Authorizer {
	return &NopAuthorizer{}
}
