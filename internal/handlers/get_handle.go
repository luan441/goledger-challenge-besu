package handlers

import (
	"encoding/json"
	"net/http"
)

type GetResponse struct {
	CurrentValue string `json:"currentValue"`
}

func GetHandle(w http.ResponseWriter, r *http.Request) {
	data := GetResponse{
		CurrentValue: "Method get current value",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
