package handlers

import (
	"encoding/json"
	"net/http"
)

type SyncResponse struct {
	Message string `json:"message"`
}

func SyncHandle(w http.ResponseWriter, r *http.Request) {
	data := SyncResponse{
		Message: "Method synchronize blockchain with database",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
