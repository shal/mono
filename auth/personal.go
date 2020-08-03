package auth

import "net/http"

type PersonalAuth struct {
	token string
}

func NewPersonal(token string) Authorizer {
	return &PersonalAuth{
		token: token,
	}
}

func (auth *PersonalAuth) Auth(r *http.Request) error {
	r.Header.Set("X-Token", auth.token)
	return nil
}

func (auth *PersonalAuth) Sign(...string) (string, error) {
	return "", nil
}