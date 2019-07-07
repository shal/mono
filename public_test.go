package mono

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Rates(t *testing.T) {
	cli := New(NewPersonalAuth("fake_token"))

	expected := []Exchange{
		{
			CodeA:     840,
			CodeB:     980,
			Date:      1552392228,
			RateSell:  27,
			RateBuy:   27.2,
			RateCross: 27.1,
		},
	}

	server := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				if err := json.NewEncoder(rw).Encode(expected); err != nil {
					t.Fail()
				}
			},
		),
	)
	defer server.Close()

	BaseURL = server.URL
	rates, err := cli.Rates()

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	for i, rate := range rates {
		if rate != expected[i] {
			t.Errorf("%v and %v is not equal", rate, expected)
		}
	}
	BaseURL = DefaultBaseURL
}
