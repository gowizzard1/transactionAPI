package utils

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}
