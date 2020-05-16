package mono

import (
	"context"
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

func TestAuthCore_GetJSON(t *testing.T) {
	core := newAuthCore(FakeAuthorizer{})

	srv, rr := FakeServer("Body", http.StatusOK)
	core.SetBaseURL(srv.URL)
	defer srv.Close()

	t.Run("makes GET request", func(t *testing.T) {
		body, status, err := core.GetJSON(context.Background(), "/", nil)
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

		body, status, err := core.GetJSON(context.Background(), "/", headers)
		rr.AssertHeaders(t, headers)

		assertEqual(t, "Body", string(body))
		assertEqual(t, http.StatusOK, status)
		assertEqual(t, nil, err)
	})

	t.Run("returns body", func(t *testing.T) {
		body, _, _ := core.GetJSON(context.Background(), srv.URL, nil)
		assertEqual(t, "Body", string(body))
	})

	t.Run("returns status", func(t *testing.T) {
		_, status, _ := core.GetJSON(context.Background(), "/", nil)
		assertEqual(t, http.StatusOK, status)
	})
}

func TestAuthCore_PostJSON(t *testing.T) {
	core := newAuthCore(FakeAuthorizer{})

	srv, rr := FakeServer("Body", http.StatusOK)
	core.SetBaseURL(srv.URL)
	defer srv.Close()

	t.Run("makes POST request", func(t *testing.T) {
		body, status, err := core.PostJSON(context.Background(), "/", nil, nil)
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

		body, status, err := core.PostJSON(context.Background(), "/", headers, nil)
		rr.AssertHeaders(t, headers)

		assertEqual(t, "Body", string(body))
		assertEqual(t, http.StatusOK, status)
		assertEqual(t, nil, err)
	})

	t.Run("returns body", func(t *testing.T) {
		body, _, _ := core.PostJSON(context.Background(), srv.URL, nil, nil)
		assertEqual(t, "Body", string(body))
	})

	t.Run("returns status", func(t *testing.T) {
		_, status, _ := core.PostJSON(context.Background(), "/", nil, nil)
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

	core.SetBaseURL(srv.URL)

	rates, err := core.Rates(context.Background())

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	for i, rate := range rates {
		if rate != expected[i] {
			t.Errorf("%v and %v is not equal", rate, expected)
		}
	}
}
