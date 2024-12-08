package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luan441/goledger-challenge-besu/internal/besu"
	"github.com/luan441/goledger-challenge-besu/internal/database"
	"github.com/luan441/goledger-challenge-besu/internal/repository"
)

type CheckResponse struct {
	Check   bool   `json:"check"`
	Message string `json:"message"`
}

func CheckHandle(w http.ResponseWriter, r *http.Request) {
	currentValue, err := besu.CallContract()
	if err != nil {
		data := CheckResponse{
			Message: fmt.Sprintf("error when processing requisition: %v", err),
			Check:   false,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(data)
		return
	}

	conn, err := database.OpenConn()
	if err != nil {
		data := CheckResponse{
			Message: fmt.Sprintf("database connection error: %v", err),
			Check:   false,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(data)
		return
	}

	bvr := repository.NewBlockchainValueRepository(conn)

	bv, err := bvr.GetLast()
	if err != nil {
		data := CheckResponse{
			Message: fmt.Sprintf("error registering the value in the database: %v", err),
			Check:   false,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(data)
		return
	}
	data := CheckResponse{
		Message: "Values successfully checked",
		Check:   bv.Value == currentValue,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
