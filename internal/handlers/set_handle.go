package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luan441/goledger-challenge-besu/internal/besu"
)

type SetResquest struct {
	Value int64 `json:"value"`
}

type SetResponse struct {
	Message string `json:"message"`
}

func SetHandle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var st SetResquest
	err := decoder.Decode(&st)
	if err != nil {
		data := SetResponse{
			Message: fmt.Sprintf("error when processing requisition: %v", err),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(data)
		return
	}

	if st.Value <= 0 {
		data := SetResponse{
			Message: "The value has not been entered or is less than or equal to 0, enter a value greater than 0",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(data)
		return
	}

	err = besu.ExecContract(st.Value)
	if err != nil {
		data := SetResponse{
			Message: fmt.Sprintf("error recording value in smart contract: %v", err),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(data)
	}

	data := SetResponse{
		Message: fmt.Sprintf("The value %v was successfuly registered", st.Value),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
