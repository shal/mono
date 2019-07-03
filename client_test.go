package mono

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetJSON(t *testing.T) {
	cli := New("fake_token")

	// Mock fake server.
	expected := "This is fake response from the fake API"
	server := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				rw.Write([]byte(expected))
			},
		),
	)
	defer server.Close()

	// Get request to the fake server to validate, that method works properly.
	BaseURL = server.URL
	body, status, err := cli.GetJSON("/fake")
	if err != nil || status != 200 || string(body) != expected {
		t.Fail()
	}
}

func TestClient_Rates(t *testing.T) {
	cli := New("fake_token")

	// Mock fake server.
	expected := []CurrencyInfo{
		{
			CurrencyCodeA: 840,
			CurrencyCodeB: 980,
			Date: 1552392228,
			RateSell: 27,
			RateBuy: 27.2,
			RateCross: 27.1,

		},
	}

	server := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				json.NewEncoder(rw).Encode(expected)
			},
		),
	)
	defer server.Close()

	// Get request to the fake server to validate, that method works properly.
	BaseURL = server.URL
	rates, err := cli.Rates()

	if err != nil {
		t.Fail()
	}

	for i, rate := range rates {
		if rate != expected[i] {
			t.Fail()
		}
	}
}