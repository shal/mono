package option

import (
	"net/http"
)

// ClientOption ...
type ClientOption interface {
	Apply(*DialSettings)
}

// DialSettings ...
type DialSettings struct {
	Endpoint   *string
	HTTPClient *http.Client
}

// ClientOptionFunc ....
type ClientOptionFunc func(*DialSettings)

// Apply ...
func (f ClientOptionFunc) Apply(settings *DialSettings) { f(settings) }

// WithEndpoint returns a ClientOption that overrides the default endpoint to be used for a service.
func WithEndpoint(url string) ClientOption {
	return ClientOptionFunc(func(settings *DialSettings) {
		settings.Endpoint = &url
	})
}

// WithHTTPClient returns a ClientOption that specifies the HTTP client to use
// as the basis of communications. This option may only be used with services
// that support HTTP as their communication transport. When used, the
// WithHTTPClient option takes precedent over all other supplied options.
func WithHTTPClient(client *http.Client) ClientOption {
	return ClientOptionFunc(func(settings *DialSettings) {
		settings.HTTPClient = client
	})
}
