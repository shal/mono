package mono

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	jsonString := []byte(`{"time": 1583000643}`)
	expected := struct{ Time Time }{
		Time: Time{time.Date(2020, 2, 29, 18, 24, 03, 00, time.UTC)},
	}

	var actual struct{ Time Time }
	err := json.Unmarshal(jsonString, &actual)

	if err != nil {
		t.Fatalf("did not expect error: %v", err)
	}

	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
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
		t.Errorf("expected error: nil, actual error: %v", err)
	}
	if expectedJson != string(actualJson) {
		t.Errorf("expected data: %v, actual data: %v", expectedJson, string(actualJson))
	}
}
