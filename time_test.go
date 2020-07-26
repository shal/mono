package mono

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	jsonString := []byte(`{"time": 1583000643}`)
	expectedData := struct{ Time Time }{
		Time: Time{time.Date(2020, 2, 29, 18, 24, 03, 00, time.UTC)},
	}

	var actualData struct{ Time Time }
	err := json.Unmarshal(jsonString, &actualData)

	if err != nil {
		t.Fatalf("expected error: nil, actual error: %v", err)
	}
	if expectedData != actualData {
		t.Fatalf("expected data: %v, actual data: %v", expectedData, actualData)
	}
}

func TestTime_MarshalJson(t *testing.T) {
	transactionData := struct {
		Time Time `json:"time"`
	}{
		Time: Time{time.Date(2020, 2, 29, 18, 24, 03, 00, time.UTC)},
	}
	expectedJson := `{"time":1583000643}`

	actualJson, err := json.Marshal(&transactionData)

	if err != nil {
		t.Fatalf("expected error: nil, actual error: %v", err)
	}
	if expectedJson != string(actualJson) {
		t.Fatalf("expected data: %v, actual data: %v", expectedJson, string(actualJson))
	}
}
