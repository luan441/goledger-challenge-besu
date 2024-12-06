package handlers

import (
	"encoding/json"
	"net/http"
)

type SetResponse struct {
	Message string `json:"message"`
}

func SetHandle(w http.ResponseWriter, r *http.Request) {
	data := SetResponse{
		Message: "Method set new value",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
