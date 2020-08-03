package option_test

import (
	"net/http"
	"testing"

	"shal.dev/mono/option"
)

func TestWithEndpoint(t *testing.T) {

	option.WithEndpoint("https://example.com")
}

func TestWithHTTPClient(t *testing.T) {

	option.WithHTTPClient(http.DefaultClient)
}
