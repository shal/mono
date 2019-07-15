package mono

import (
	"testing"
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

//func TestPersonal_Transactions(t *testing.T) {
//	transactions := []Transaction{
//		{
//			ID:              "ZuHWzqkKGVo=",
//			Time:            1554466347,
//			Description:     "Покупка щастя",
//			MCC:             7997,
//			Hold:            false,
//			Amount:          -95000,
//			OperationAmount: -95000,
//			CurrencyCode:    980,
//			CommissionRate:  0,
//			CashBackAmount:  19000,
//		},
//	}
//
//}

//func TestPersonal_SetWebHook(t *testing.T) {
//
//}
