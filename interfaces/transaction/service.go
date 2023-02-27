package transaction

type IService interface {
	Debit(amount float64) error
	Credit(amount float64) error
	Balance() float64
}
