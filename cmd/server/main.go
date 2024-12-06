package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luan441/goledger-challenge-besu/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/blockchain/set", handlers.SetHandle)
	mux.HandleFunc("GET /api/blockchain/get", handlers.GetHandle)
	mux.HandleFunc("GET /api/blockchain/sync", handlers.SyncHandle)
	mux.HandleFunc("GET /api/blockchain/check", handlers.CheckHandle)

	fmt.Println("Server starting at URL http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
