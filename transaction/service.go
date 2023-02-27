package transaction

import (
	"sync"
	"transationAPI/interfaces/transaction"
)

type service struct {
	balance float64
	mux     sync.Mutex
	repo    Repository
}

func NewService(balance float64, mux *sync.Mutex, repo Repository) transaction.IService {
	return &service{
		balance: balance,
		mux:     *mux,
		repo:    repo,
	}
}

func (s *service) Debit(amount float64) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.balance -= amount
	return nil
}

func (s *service) Credit(amount float64) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.balance += amount
	return nil
}

func (s *service) Balance() float64 {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.balance
}
