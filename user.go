package mono

// Overview of user and related accounts.
type UserInfo struct {
	// User name.
	Name string `json:"name"`
	// URL for receiving new transactions.
	WebHookURL string `json:"webHookUrl"`
	// List of available accounts
	Accounts []Account `json:"accounts"`
}
