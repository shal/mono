package mono

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_Statement(t *testing.T) {
	cli := New(NewPersonalAuth("fake_token"))

	expected := []StatementItem{
		{
			ID:              "ZuHWzqkKGVo=",
			Time:            1554466347,
			Description:     "Покупка щастя",
			MCC:             7997,
			Hold:            false,
			Amount:          -95000,
			OperationAmount: -95000,
			CurrencyCode:    980,
			CommissionRate:  0,
			CashbackAmount:  19000,
			Balance:         10050000,
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

	from := time.Now().AddDate(0, 0, -10)
	to := time.Now()

	items, err := cli.Statement("0", from, to)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	for i, item := range items {
		if item != expected[i] {
			t.Errorf("%v and %v is not equal", item, expected)
		}
	}
}

func TestClient_User(t *testing.T) {
	cli := New(NewPersonalAuth("fake_token"))

	expected := UserInfo{
		Name:       "John Doe",
		WebHookURL: "http://localhost:8080",
		Accounts: []Account{
			{
				ID:           "kKGVoZuHWzqVoZuH",
				Balance:      10000000,
				CreditLimit:  10000000,
				CurrencyCode: 980,
				CashbackType: None,
			},
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

	user, err := cli.User()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if user == &expected {
		t.Errorf("%v and %v is not equal", *user, expected)
	}
}

func TestClient_SetWebHook(t *testing.T) {
	cli := New(NewPersonalAuth("fake_token"))

	server := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				if _, err := rw.Write([]byte(`{ "status": "ok" }`)); err != nil {
					t.Fail()
				}
			},
		),
	)
	defer server.Close()
	BaseURL = server.URL

	body, err := cli.SetWebHook("http://127.0.0.1:8080")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	var payload map[string]string
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if payload["status"] != "ok" {
		t.Errorf("%v and ok is not equal", payload["status"])
	}
}
