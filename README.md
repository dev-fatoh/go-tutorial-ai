# Go Bank Web UI

A small Go application that serves a simple banking dashboard in the browser.

## Overview

This project now includes:

- a web form for deposits and withdrawals,
- a styled balance panel,
- a Go HTTP server that serves the frontend and handles API requests.

## Features

- View the current account balance
- Deposit money through a form
- Withdraw money through a form
- Handle invalid amounts and insufficient funds with messages
- Serve the UI and API from the same Go program

## Project Files

- `main.go` - Starts the server and handles `/api/balance` and `/api/transaction`
- `index.html` - Page layout for the banking form
- `styles.css` - Styling for the dashboard
- `script.js` - Frontend logic for loading balance and submitting transactions

## How to Run

1. Make sure Go is installed.
2. From the project folder, run:

   ```bash
   go run .
   ```

3. Open your browser and visit:

   ```text
   http://localhost:8080
   ```

## API Endpoints

- `GET /api/balance` - Returns the current balance as JSON
- `POST /api/transaction` - Accepts a JSON body with `action` and `amount`

Example request body:

```json
{
  "action": "deposit",
  "amount": 50
}
```

## Notes

- The balance starts at `1000.00` when the server starts.
- The balance is stored in memory while the program is running.
- This project is a beginner-friendly example of combining Go server code with a simple frontend.
