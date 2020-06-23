package upcloud

// Account represents an UpCloud account
type Account struct {
	// UpCloud account username
	Username string `json:"username"`
	// UpCloud account credits
	Credits string `json:"credits"`
}

// getAccountResponse is a response wrapper to match the UpCloud API payload
type getAccountResponse struct {
	Account *Account `json:"account"`
}
