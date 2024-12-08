package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luan441/goledger-challenge-besu/internal/besu"
	"github.com/luan441/goledger-challenge-besu/internal/database"
	"github.com/luan441/goledger-challenge-besu/internal/repository"
)

type SyncResponse struct {
	Message string `json:"message"`
}

func SyncHandle(w http.ResponseWriter, r *http.Request) {
	currentValue, err := besu.CallContract()
	if err != nil {
		data := SyncResponse{
			Message: fmt.Sprintf("error when processing requisition: %v", err),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(data)
		return
	}

	conn, err := database.OpenConn()
	if err != nil {
		data := SyncResponse{
			Message: fmt.Sprintf("database connection error: %v", err),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(data)
		return
	}

	bvr := repository.NewBlockchainValueRepository(conn)

	_, err = bvr.Insert(currentValue)
	if err != nil {
		data := SyncResponse{
			Message: fmt.Sprintf("error registering the value in the database: %v", err),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(data)
		return
	}

	data := SyncResponse{
		Message: "synchronized database",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
