package handlers

import (
	"encoding/json"
	"net/http"
)

type CheckResponse struct {
	Check   bool   `json:"check"`
	Message string `json:"message"`
}

func CheckHandle(w http.ResponseWriter, r *http.Request) {
	data := CheckResponse{
		Message: "Method check current value blockchain with database",
		Check:   true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
