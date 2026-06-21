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

## What's Changed (Inline Field Errors)

- Added per-field inline error messaging to the form inputs so users see validation feedback next to the related field instead of only a global message.
- Client-side: `script.js` now validates input and shows/clears field errors (`amount-error`, `action-error`). It also displays any field-specific errors returned by the server.
- Server-side: `main.go` now returns structured validation responses with an `errors` object mapping field names to messages (e.g. `{ "errors": { "amount": "Insufficient funds" }, "error": "Insufficient funds" }`).
- Accessibility: inputs include `aria-describedby` and `aria-live` attributes for polite announcement of inline errors.

## Files Modified

- [index.html](index.html): Added inline error elements and `aria-describedby` attributes for inputs.
- [script.js](script.js): Added `setFieldError`/`clearFieldErrors`, client-side validation improvements, and handling for server-returned `errors`.
- [main.go](main.go): Added structured `errors` to JSON responses for validation failures.
- [styles.css](styles.css): Added `.input-error` and `.field-error` styles.

## Testing the Changes

1. Start the server from the project root:

```bash
go run main.go
```

2. Open `http://localhost:8080` and try these scenarios:

- Submit an empty or non-positive `Amount` → inline error appears under the `Amount` field.
- Try to withdraw more than the current balance → server returns a field error and it appears inline under `Amount`.
- Change the transaction type to an invalid value (if testing API directly) → server returns an `action` field error.

3. Alternatively, test the API directly with `curl`:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"action":"withdraw","amount":99999}' http://localhost:8080/api/transaction
```

The response will include an `errors` object when validation fails.
