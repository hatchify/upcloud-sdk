package upcloud

// Account represents an UpCloud account
type Account struct {
	Credits  string `json:"credits"`
	Username string `json:"username"`
}

type getAccountResponse struct {
	Account *Account `json:"account"`
}
