package mono

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakeAuthorizer struct{}

func (fake FakeAuthorizer) Auth(request *http.Request) error {
	request.Header.Set("X-Auth-Test", "Success")
	return nil
}

func fix() {
	BaseURL = DefaultBaseURL
}

func TestAuthCore_GetJSON(t *testing.T) {
	core := newAuthCore(FakeAuthorizer{})

	srv, rr := FakeServer("Body", http.StatusOK)
	BaseURL = srv.URL
	defer srv.Close()
	defer fix()

	t.Run("makes GET request", func(t *testing.T) {
		body, status, err := core.GetJSON("/", nil)
		rr.AssertMethod(t, "GET")

		assertEqual(t, "Body", string(body))
		assertEqual(t, http.StatusOK, status)
		assertEqual(t, nil, err)
	})

	t.Run("handles headers", func(t *testing.T) {
		headers := map[string]string{
			"X-Test1":     "Test",
			"X-Test2":     "Test",
			"X-Auth-Test": "Success",
		}

		body, status, err := core.GetJSON("/", headers)
		rr.AssertHeaders(t, headers)

		assertEqual(t, "Body", string(body))
		assertEqual(t, http.StatusOK, status)
		assertEqual(t, nil, err)
	})

	t.Run("returns body", func(t *testing.T) {
		body, _, _ := core.GetJSON(srv.URL, nil)
		assertEqual(t, "Body", string(body))
	})

	t.Run("returns status", func(t *testing.T) {
		_, status, _ := core.GetJSON("/", nil)
		assertEqual(t, http.StatusOK, status)
	})
}

func TestAuthCore_PostJSON(t *testing.T) {
	core := newAuthCore(FakeAuthorizer{})

	srv, rr := FakeServer("Body", http.StatusOK)
	BaseURL = srv.URL
	defer srv.Close()
	defer fix()

	t.Run("makes POST request", func(t *testing.T) {
		body, status, err := core.PostJSON("/", nil, nil)
		rr.AssertMethod(t, "POST")

		assertEqual(t, "Body", string(body))
		assertEqual(t, http.StatusOK, status)
		assertEqual(t, nil, err)
	})

	t.Run("handles headers", func(t *testing.T) {
		headers := map[string]string{
			"X-Test1":     "Test",
			"X-Test2":     "Test",
			"X-Auth-Test": "Success",
		}

		body, status, err := core.PostJSON("/", headers, nil)
		rr.AssertHeaders(t, headers)

		assertEqual(t, "Body", string(body))
		assertEqual(t, http.StatusOK, status)
		assertEqual(t, nil, err)
	})

	t.Run("returns body", func(t *testing.T) {
		body, _, _ := core.PostJSON(srv.URL, nil, nil)
		assertEqual(t, "Body", string(body))
	})

	t.Run("returns status", func(t *testing.T) {
		_, status, _ := core.PostJSON("/", nil, nil)
		assertEqual(t, http.StatusOK, status)
	})
}

func TestAuthCore_Rates(t *testing.T) {
	core := newCore()

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

	srv := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				if err := json.NewEncoder(rw).Encode(expected); err != nil {
					t.Fail()
				}
			},
		),
	)
	defer srv.Close()

	BaseURL = srv.URL
	rates, err := core.Rates()

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
