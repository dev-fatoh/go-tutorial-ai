// Package main implements a tiny HTTP server that serves a simple
// banking UI and provides two API endpoints for balance and transactions.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// accountBalance holds the in-memory balance for this toy example.
var accountBalance = 1000.00

// transactionRequest represents the JSON payload expected for transactions.
type transactionRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

// response is the generic API response shape. `Errors` maps field names
// to messages for structured validation feedback.
type response struct {
	Balance float64 `json:"balance,omitempty"`
	Message string  `json:"message,omitempty"`
	Error   string  `json:"error,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func main() {
	// Route handlers for API and static files.
	http.HandleFunc("/api/balance", handleBalance)
	http.HandleFunc("/api/transaction", handleTransaction)
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleBalance(w http.ResponseWriter, r *http.Request) {
	// Only allow GET for balance retrieval.
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Balance: accountBalance})
}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	// Only accept POST for transactions.
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode request body into the typed struct.
	var req transactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{Error: "Invalid request payload"})
		return
	}

	// Basic server-side validation for a positive amount.
	if req.Amount <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{Errors: map[string]string{"amount": "Amount must be greater than zero"}, Error: "Validation failed"})
		return
	}

	// Handle the supported transaction types.
	switch req.Action {
	case "deposit":
		accountBalance += req.Amount
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response{
			Balance: accountBalance,
			Message: fmt.Sprintf("Deposited $%.2f successfully", req.Amount),
		})
	case "withdraw":
		// Prevent overdraft in this simple example.
		if req.Amount > accountBalance {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response{Errors: map[string]string{"amount": "Insufficient funds"}, Error: "Insufficient funds"})
			return
		}
		accountBalance -= req.Amount
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response{
			Balance: accountBalance,
			Message: fmt.Sprintf("Withdrew $%.2f successfully", req.Amount),
		})
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{Errors: map[string]string{"action": "Invalid transaction type"}, Error: "Invalid transaction type"})
	}
}
