package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleBalance_GET(t *testing.T) {
    // ensure deterministic starting balance
    accountBalance = 500.00

    req := httptest.NewRequest(http.MethodGet, "/api/balance", nil)
    rr := httptest.NewRecorder()

    handleBalance(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %d", rr.Code)
    }

    var resp response
    if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
        t.Fatalf("decoding response: %v", err)
    }

    if resp.Balance != accountBalance {
        t.Fatalf("expected balance %v, got %v", accountBalance, resp.Balance)
    }
}

func TestHandleTransaction_DepositWithdrawAndErrors(t *testing.T) {
    // deposit
    accountBalance = 1000.00
    dep := transactionRequest{Action: "deposit", Amount: 50}
    b, _ := json.Marshal(dep)
    req := httptest.NewRequest(http.MethodPost, "/api/transaction", bytes.NewReader(b))
    rr := httptest.NewRecorder()
    handleTransaction(rr, req)
    if rr.Code != http.StatusOK {
        t.Fatalf("deposit expected 200, got %d", rr.Code)
    }
    var r1 response
    if err := json.NewDecoder(rr.Body).Decode(&r1); err != nil {
        t.Fatalf("decoding deposit response: %v", err)
    }
    if r1.Balance != accountBalance {
        t.Fatalf("after deposit expected balance %v, got %v", accountBalance, r1.Balance)
    }

    // withdraw success
    accountBalance = 200.00
    wid := transactionRequest{Action: "withdraw", Amount: 50}
    b, _ = json.Marshal(wid)
    req = httptest.NewRequest(http.MethodPost, "/api/transaction", bytes.NewReader(b))
    rr = httptest.NewRecorder()
    handleTransaction(rr, req)
    if rr.Code != http.StatusOK {
        t.Fatalf("withdraw expected 200, got %d", rr.Code)
    }
    var r2 response
    if err := json.NewDecoder(rr.Body).Decode(&r2); err != nil {
        t.Fatalf("decoding withdraw response: %v", err)
    }
    if r2.Balance != accountBalance {
        t.Fatalf("after withdraw expected balance %v, got %v", accountBalance, r2.Balance)
    }

    // withdraw insufficient funds
    accountBalance = 10.00
    wid2 := transactionRequest{Action: "withdraw", Amount: 50}
    b, _ = json.Marshal(wid2)
    req = httptest.NewRequest(http.MethodPost, "/api/transaction", bytes.NewReader(b))
    rr = httptest.NewRecorder()
    handleTransaction(rr, req)
    if rr.Code != http.StatusBadRequest {
        t.Fatalf("expected 400 for insufficient funds, got %d", rr.Code)
    }
    var r3 response
    if err := json.NewDecoder(rr.Body).Decode(&r3); err != nil {
        t.Fatalf("decoding insufficient funds response: %v", err)
    }
    if msg, ok := r3.Errors["amount"]; !ok || msg == "" {
        t.Fatalf("expected amount field error, got %v", r3.Errors)
    }

    // invalid amount (non-positive)
    accountBalance = 100.00
    bad := transactionRequest{Action: "deposit", Amount: 0}
    b, _ = json.Marshal(bad)
    req = httptest.NewRequest(http.MethodPost, "/api/transaction", bytes.NewReader(b))
    rr = httptest.NewRecorder()
    handleTransaction(rr, req)
    if rr.Code != http.StatusBadRequest {
        t.Fatalf("expected 400 for invalid amount, got %d", rr.Code)
    }
    var r4 response
    if err := json.NewDecoder(rr.Body).Decode(&r4); err != nil {
        t.Fatalf("decoding invalid amount response: %v", err)
    }
    if _, ok := r4.Errors["amount"]; !ok {
        t.Fatalf("expected amount validation error, got %v", r4.Errors)
    }

    // invalid action
    accountBalance = 100.00
    inv := transactionRequest{Action: "invalid", Amount: 10}
    b, _ = json.Marshal(inv)
    req = httptest.NewRequest(http.MethodPost, "/api/transaction", bytes.NewReader(b))
    rr = httptest.NewRecorder()
    handleTransaction(rr, req)
    if rr.Code != http.StatusBadRequest {
        t.Fatalf("expected 400 for invalid action, got %d", rr.Code)
    }
    var r5 response
    if err := json.NewDecoder(rr.Body).Decode(&r5); err != nil {
        t.Fatalf("decoding invalid action response: %v", err)
    }
    if _, ok := r5.Errors["action"]; !ok {
        t.Fatalf("expected action validation error, got %v", r5.Errors)
    }
}

func TestMethodNotAllowed(t *testing.T) {
    // GET on transaction should be method not allowed
    req := httptest.NewRequest(http.MethodGet, "/api/transaction", nil)
    rr := httptest.NewRecorder()
    handleTransaction(rr, req)
    if rr.Code != http.StatusMethodNotAllowed {
        t.Fatalf("expected 405 for GET on /api/transaction, got %d", rr.Code)
    }

    // POST on balance should be method not allowed
    req = httptest.NewRequest(http.MethodPost, "/api/balance", nil)
    rr = httptest.NewRecorder()
    handleBalance(rr, req)
    if rr.Code != http.StatusMethodNotAllowed {
        t.Fatalf("expected 405 for POST on /api/balance, got %d", rr.Code)
    }
}
