package mono

// CashBackType is type of cash-back that credits to the account.
type CashBackType string

const (
	None  = "None"  // None is cash-back type for foreign currencies accounts.
	UAH   = "UAH"   // UAH is default cash-back type for almost all accounts.
	Miles = "Miles" // Miles available only on Iron Card.
)

// Overview of user and related accounts.
type UserInfo struct {
	Name       string    `json:"name"`       // User name.
	WebHookURL string    `json:"webHookUrl"` // URL for receiving new transactions.
	Accounts   []Account `json:"accounts"`   // List of available accounts
}

// TokenRequest is representation of payload, received on corporate auth.
type TokenRequest struct {
	TokenRequestID string `json:"tokenRequestId"` // Unique token request ID.
	AcceptURL      string `json:"acceptUrl"`      // URL to redirect client or build QR on top of it.
}

// Account is simple representation of bank account.
type Account struct {
	ID           string       `json:"id"`           // Account identifier.
	Balance      int          `json:"balance"`      // Balance is minimal units (cents).
	CreditLimit  int          `json:"creditLimit"`  // Credit limit.
	CurrencyCode int          `json:"currencyCode"` // Currency code in ISO4217.
	CashBackType CashBackType `json:"cashbackType"` // Type of cash-back.
}

// Transaction is a banking transaction.
type Transaction struct {
	ID              string `json:"id"`              // Unique transaction ID.
	Time            int32  `json:"time"`            // Unix time of transaction.
	Description     string `json:"description"`     // Message attached to transaction.
	MCC             int32  `json:"mcc"`             // Merchant Category Code using ISO18245.
	Hold            bool   `json:"hold"`            // Authorization hold.
	Amount          int64  `json:"amount"`          // Amount in account currency (cents).
	OperationAmount int64  `json:"operationAmount"` // Amount in transaction currency (cents).
	CurrencyCode    int32  `json:"currencyCode"`    // Currency code using ISO4217.
	CommissionRate  int64  `json:"commissionRate"`  // Amount of commission in account currency.
	CashBackAmount  int64  `json:"cashbackAmount"`  // Amount of cash-back in account currency.
	Balance         int64  `json:"balance"`         // Balance in account currency.
}
