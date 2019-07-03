package mono

// CashbackType is type of cashback that credits to the account.
type CashbackType string

const (
	None  = "None"
	UAH   = "UAH"
	Miles = "Miles"
)

type Account struct {
	// Account identifier.
	ID string `json:"id"`
	// Balance is minimal units (cents).
	Balance int `json:"balance"`
	// Credit limit.
	CreditLimit int `json:"creditLimit"`
	// Currency code in ISO4217.
	CurrencyCode int `json:"currencyCode"`
	// Type of cashback.
	CashbackType CashbackType `json:"cashbackType"`
}

type StatementItem struct {
	ID              string `json:"id"`
	Time            int32  `json:"time"`
	Description     string `json:"description"`
	MCC             int32  `json:"mcc"`
	Hold            bool   `json:"hold"`
	Amount          int64  `json:"amount"`
	OperationAmount int64  `json:"operationAmount"`
	CurrencyCode    int32  `json:"currencyCode"`
	CommissionRate  int64  `json:"commissionRate"`
	CashbackAmount  int64  `json:"cashbackAmount"`
	Balance         int64  `json:"balance"`
}
