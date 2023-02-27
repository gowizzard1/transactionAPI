package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"transationAPI/interfaces/transaction"
	"transationAPI/utils"
)

type Controller struct {
	service transaction.IService
}

func NewController(service transaction.IService) Controller {
	return Controller{
		service: service,
	}
}

func (c Controller) Balance(w http.ResponseWriter, r *http.Request) {
	//call service
	bal := c.service.Balance()
	resp := transaction.BalanceResponse{
		Balance: *bal,
	}
	if err := utils.Response(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request) {
	req := transaction.TransactionRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call service
	err = c.service.CreateTransaction(req)
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
