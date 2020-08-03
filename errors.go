package mono

import "errors"

var (
	ErrUnexpectedAuth = errors.New("unexpected authorizer")
)

// Error is a simple representation of MonoBank API error.
type Error struct {
	ErrorDescription string `json:"errorDescription"`
}

func (e Error) Error() string {
	return e.ErrorDescription
}
