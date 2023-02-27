package transaction

type IService interface {
	CreateTransaction(trans TransactionRequest) error
	Balance() *float64
}
