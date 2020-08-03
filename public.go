package mono

import (
	"shal.dev/mono/auth"
	"shal.dev/mono/option"
)

// Public gives access to public methods.
type Public struct {
	*Client
}

// NewPublic returns new client of MonoBank Public API.
func NewPublic(opts ...option.ClientOption) (*Public, error) {
	client, err := NewClient(auth.NopAuth(), opts...)
	if err != nil {
		return nil, err
	}

	return &Public{
		Client: client,
	}, nil
}
