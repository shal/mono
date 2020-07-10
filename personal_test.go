package mono

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestNewPersonal(t *testing.T) {
	personal := NewPersonal("token")

	if personal == nil {
		t.Error()
	}
}

func TestPersonal_User(t *testing.T) {

}

//func TestPersonal_User(t *testing.T) {
//	user := UserInfo{
//		Name:       "John Doe",
//		WebHookURL: "https://localhost:8080/test/",
//		Accounts: []Account{
//			{
//				ID:           "kKGVoZuHWzqVoZuH",
//				Balance:      10000000,
//				CreditLimit:  10000000,
//				CurrencyCode: 980,
//				CashBackType: "UAH",
//			},
//		},
//	}

//client := NewPersonal("fake")

//srv, rr := FakeServer("Test", http.StatusOK)
//BaseURL = srv.URL
//defer srv.Close()
//resp, err := client.User()

//}

func TestPersonal_Transactions(t *testing.T) {
	client := NewPersonal("Success")
	mockResponse := `
[
    {
        "id": "zxcvtyuioasdfghj",
        "time": 1593259442,
        "description": "5************6",
        "mcc": 4829,
        "amount": -128200,
        "operationAmount": -128200,
        "currencyCode": 980,
        "commissionRate": 0,
        "cashbackAmount": 0,
        "balance": 0,
        "hold": true
    },
    {
        "id": "qwertyuioasdfghj",
        "time": 1583000643,
        "description": "Киевстар\n+380688888888",
        "mcc": 4814,
        "amount": -20000,
        "operationAmount": -20000,
        "currencyCode": 980,
        "commissionRate": 0,
        "cashbackAmount": 0,
        "balance": 128200,
        "hold": true
    }
]`
	expectedTransactions := []Transaction{
		{
			ID:              "zxcvtyuioasdfghj",
			Time:            Time{time.Date(2020, 6, 27, 12, 04, 02, 00, time.UTC)},
			Description:     "5************6",
			MCC:             4829,
			Hold:            true,
			Amount:          -128200,
			OperationAmount: -128200,
			CurrencyCode:    980,
			CommissionRate:  0,
			CashBackAmount:  0,
			Balance:         0,
		},
		{
			ID:              "qwertyuioasdfghj",
			Time:            Time{time.Date(2020, 2, 29, 18, 24, 03, 00, time.UTC)},
			Description:     "Киевстар\n+380688888888",
			MCC:             4814,
			Hold:            true,
			Amount:          -20000,
			OperationAmount: -20000,
			CurrencyCode:    980,
			CommissionRate:  0,
			CashBackAmount:  0,
			Balance:         128200,
		},
	}
	srv, _ := FakeServer(mockResponse, http.StatusOK)
	defer srv.Close()
	client.SetBaseURL(srv.URL)

	actual, err := client.Transactions(context.TODO(), "acc", time.Now(), time.Now())

	if err != nil {
		t.Fatalf("expected error: nil, actual error: %v", err)
	}
	for i, v := range expectedTransactions {
		if actual[i] != v {
			t.Fatalf("expected transactions[%v]: %v, actual transactions[%v]: %v", i, expectedTransactions[i], i, actual[i])
		}
	}
}

//func TestPersonal_SetWebHook(t *testing.T) {
//
//}
