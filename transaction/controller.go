package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"transationAPI/interfaces/transaction"
	"transationAPI/utils"
)

type transactionRequest struct {
	Amount float64 `json:"amount"`
}

type Controller struct {
	service transaction.IService
}

func NewController(service transaction.IService) Controller {
	return Controller{
		service: service,
	}
}

func (c Controller) Debit(w http.ResponseWriter, r *http.Request) {
	req := transactionRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call service
	c.service.Debit(req.Amount)
	status := "success"
	message := fmt.Sprintf("Debit of $%.2f successful.", req.Amount)
	if err != nil {
		status = "error"
		message = err.Error()
	}
	resp := transaction.TransactionResponse{
		Status:  status,
		Message: message,
	}
	if err := utils.Response(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c Controller) Credit(w http.ResponseWriter, r *http.Request) {
	req := transactionRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call service
	err = c.service.Credit(req.Amount)
	status := "success"
	message := fmt.Sprintf("Credit of $%.2f successful.", req.Amount)
	if err != nil {
		status = "error"
		message = err.Error()
	}
	resp := transaction.TransactionResponse{
		Status:  status,
		Message: message,
	}
	if err := utils.Response(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c Controller) Balance(w http.ResponseWriter, r *http.Request) {
	//call service
	bal := c.service.Balance()
	resp := transaction.BalanceResponse{
		Balance: bal,
	}
	if err := utils.Response(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
