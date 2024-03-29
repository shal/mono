package mono

// CashBackType is type of cash-back that credits to the account.
type CashBackType string

const (
	// None is cash-back type for foreign currencies accounts.
	None CashBackType = "None"
	// UAH is default cash-back type for almost all accounts.
	UAH CashBackType = "UAH"
	// Miles available only on Iron Card.
	Miles CashBackType = "Miles"
)

// AccountType is type of the account.
type AccountType string

const (
	Platinum AccountType = "platinum"
	White    AccountType = "white"
	Black    AccountType = "black"
	FOP      AccountType = "fop"
	EAid     AccountType = "eAid" // єПідтримка

)

// UserInfo is an overview of user and related accounts.
type UserInfo struct {
	ID         string    `json:"clientId"`
	Name       string    `json:"name"`       // User name.
	WebHookURL string    `json:"webHookUrl"` // URL for receiving new transactions.
	Accounts   []Account `json:"accounts"`   // List of available accounts
	Jars       []Jar     `json:"jars"`
}

type Jar struct {
	ID           string `json:"id"`
	SendID       string `json:"sendId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CurrencyCode int    `json:"currencyCode"`
	Balance      int64  `json:"balance"`
	Goal         int64  `json:"goal"`
}

// TokenRequest is representation of payload, received on corporate auth.
type TokenRequest struct {
	TokenRequestID string `json:"tokenRequestId"` // Unique token request ID.
	AcceptURL      string `json:"acceptUrl"`      // URL to redirect client or build QR on top of it.
}

// Account is simple representation of bank account.
type Account struct {
	ID           string       `json:"id"` // Account identifier.
	SendID       string       `json:"sendId"`
	Balance      int          `json:"balance"`      // Balance is minimal units (cents).
	CreditLimit  int          `json:"creditLimit"`  // Credit limit.
	CurrencyCode int32        `json:"currencyCode"` // Currency code in ISO4217.
	CashBackType CashBackType `json:"cashbackType"` // Type of cash-back.
	Type         AccountType  `json:"type"`         // Type of card.
	IBAN         string       `json:"iban"`         // IBAN.
	MaskedPan    []string     `json:"maskedPan"`
}

// Transaction is a banking transaction.
type Transaction struct {
	ID              string `json:"id"`          // Unique transaction ID.
	Time            Time   `json:"time"`        // UTC time of transaction.
	Description     string `json:"description"` // Message attached to transaction.
	MCC             int32  `json:"mcc"`         // Merchant Category Code using ISO18245.
	OriginalMCC     int32  `json:"originalMcc"`
	Hold            bool   `json:"hold"`            // Authorization hold.
	Amount          int64  `json:"amount"`          // Amount in account currency (cents).
	OperationAmount int64  `json:"operationAmount"` // Amount in transaction currency (cents).
	CurrencyCode    int32  `json:"currencyCode"`    // Currency code using ISO4217.
	CommissionRate  int64  `json:"commissionRate"`  // Amount of commission in account currency.
	CashBackAmount  int64  `json:"cashbackAmount"`  // Amount of cash-back in account currency.
	Balance         int64  `json:"balance"`         // Balance in account currency.
	Comment         string `json:"comment"`
	ReceiptID       string `json:"receiptId"` // ID of the receipt.
	InvoiceID       string `json:"invoiceId"` // ID of the invoice.
	EDRPOU          string `json:"counterEdrpou"`
	IBAN            string `json:"counterIban"`
}
