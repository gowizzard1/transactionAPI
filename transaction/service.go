package transaction

import (
	"transationAPI/interfaces/transaction"
	"transationAPI/models"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) transaction.IService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTransaction(trans transaction.TransactionRequest) error {
	// get current balance
	bal, err := s.repo.GetBalance()
	if err != nil {
		// handle error
	}
	var amount float64
	switch trans.Type {
	case "credit":
		amount = trans.Amount

	case "debit":
		amount = -(trans.Amount)
	}
	newBal := *bal + amount
	transaction := models.Transaction{
		Amount:  trans.Amount,
		Type:    trans.Type,
		Balance: newBal,
	}
	if err := s.repo.Create(&transaction); err != nil {
		return err
	}
	return nil
}

func (s *service) Balance() *float64 {
	bal, err := s.repo.GetBalance()
	if err != nil {
		// handle error
	}
	return bal
}
