package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luan441/goledger-challenge-besu/internal/besu"
)

type GetResponse struct {
	CurrentValue int64  `json:"currentValue"`
	Message      string `json:"message"`
}

func GetHandle(w http.ResponseWriter, r *http.Request) {
	currentValue, err := besu.CallContract()
	if err != nil {
		data := GetResponse{
			Message:      fmt.Sprintf("error when processing requisition: %v", err),
			CurrentValue: 0,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(data)
		return
	}

	data := GetResponse{
		CurrentValue: currentValue,
		Message:      "Value successfuly recovered",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
