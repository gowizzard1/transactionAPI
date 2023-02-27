package transaction

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

type TransactionResponse struct {
	Status  string
	Message string
}

type TransactionRequest struct {
	Amount      float64
	Type        string
	Description string
}
