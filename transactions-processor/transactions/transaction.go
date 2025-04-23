package transactions

type Transaction struct {
	TransactionID string `json:"transactionId"`
	Amount        int    `json:"amount"`
	UserId        string `json:"userId"`
	AccountID     string `json:"accountId"`
	Type          string `json:"type"`
	Method        string `json:"method"`
	Destination   string `json:"destination"`
}