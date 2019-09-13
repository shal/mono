package mono

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type RequestRecorder struct {
	Headers http.Header
	Method  string
	Payload []byte
}

func FakeServer(text string, code int) (*httptest.Server, *RequestRecorder) {
	rr := new(RequestRecorder)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rr.Headers = r.Header
		rr.Method = r.Method
		rr.Payload, _ = ioutil.ReadAll(r.Body)

		w.WriteHeader(code)

		_, err := w.Write([]byte(text))
		if err != nil {
			panic(err)
		}
	})

	return httptest.NewServer(handler), rr
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func (rr *RequestRecorder) AssertHeaders(t *testing.T, expected map[string]string) {
	for k, v := range expected {
		assertEqual(t, v, rr.Headers.Get(k))
	}
}

func (rr *RequestRecorder) AssertMethod(t *testing.T, expected string) {
	assertEqual(t, expected, rr.Method)
}

func (rr *RequestRecorder) AssertBody(t *testing.T, expected []byte) {
	assertEqual(t, expected, rr.Payload)
}

func TestCore_GetJSON(t *testing.T) {
	core := newCore()

	srv, rr := FakeServer("Body", http.StatusOK)
	BaseURL = srv.URL
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
			"X-Test1": "Test",
			"X-Test2": "Test",
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

func TestCore_PostJSON(t *testing.T) {
	core := newCore()

	srv, rr := FakeServer("Body", http.StatusOK)
	BaseURL = srv.URL
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
			"X-Test1": "Test",
			"X-Test2": "Test",
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

func TestCore_Rates(t *testing.T) {
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
	rates, err := core.Rates(context.Background())

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
