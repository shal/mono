package mono

import "testing"

func TestCurrencyFromISO4217(t *testing.T) {
	for code, expected := range currencyCodes {
		ccy, err := CurrencyFromISO4217(code)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		if ccy != expected {
			t.Errorf("%v and %v is not equal", ccy, expected)
		}
	}
}