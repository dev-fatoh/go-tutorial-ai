# Copilot Instructions & Prompts — Go Bank Web UI

Purpose

- Central place for reusable prompts, task plans, and guidelines to help Copilot-style agents work on this repo.

How to use

- Copy a prompt below into the chat to start a task. Use the plan templates as checklists when implementing features or fixes.

---

## Prompts

1. Add feature: Inline validation improvements
   """
   Implement improved client-side and server-side validation for the transaction form in this repository. Provide per-field inline errors, ensure ARIA attributes are set for accessibility, and return structured `errors` in JSON from the server. Update styles and README. Run the app locally and verify behavior.
   """

2. Fix bug: Invalid amount handling
   """
   Locate the code handling transaction amounts. If a negative or non-numeric amount is submitted, make sure the frontend displays an inline error and the server returns a structured validation error. Add unit tests for `handleTransaction` around invalid amounts.
   """

3. Add tests: handleTransaction
   """
   Add table-driven unit tests for `handleTransaction` covering deposit, withdraw, insufficient funds, invalid payload, and invalid transaction type. Use the standard library `net/http/httptest` and assert responses include `errors` where applicable.
   """

4. Accessibility audit
   """
   Audit the UI for accessibility issues (labels, ARIA live regions, focus management). Propose and implement fixes such as `aria-describedby`, `aria-invalid`, and focusing the first invalid field after form submit.
   """

5. Create commit and push
   """
   Stage all changes with `git add -A`, commit with a clear message, and push to a remote branch. If no remote exists, add instructions to create one.
   """

6. Refactor: API handlers
   """
   Refactor `main.go` to move API handlers into `api/handlers.go` and add a simple `server.Start()` entry in `main.go`. Keep behavior identical and add minimal tests for handler functions.
   """

7. Add pre-commit hooks
   """
   Add a `.git/hooks/pre-commit` (or use `husky`/`lefthook`) to run `gofmt`, `go vet`, and lint checks, and to ensure staged JS/CSS/HTML are formatted.
   """

---

## Task Plans (Templates)

A) Plan: Add field-level inline errors

1. Add inline error containers next to relevant inputs in `index.html`.
2. Add CSS classes for `.input-error` and `.field-error` in `styles.css`.
3. Implement `setFieldError` and `clearFieldErrors` helpers in `script.js`.
4. Add client-side validation and show errors inline before submitting.
5. Update server to return structured `errors` map for validation failures.
6. Update `README.md` with changes and testing steps.
7. Run manual tests and adjust UI/ARIA as needed.

B) Plan: Add unit tests for `handleTransaction`

1. Create `main_test.go` or `api/handlers_test.go`.
2. Use `httptest.NewRecorder()` and table-driven tests.
3. Check HTTP status codes and JSON bodies, especially `errors` mapping.
4. Run `go test ./...` and fix failures.

C) Plan: Accessibility improvements

1. Ensure each input has a visible label and `aria-describedby` for errors.
2. Use `aria-live="polite"` for inline error elements.
3. After validation failure, set focus to the first invalid input.
4. Run Lighthouse accessibility audit and address major issues.

---

## Commit Message Templates

- Feature: "feat: add inline field errors and structured validation"
- Fix: "fix: validate transaction amount and show inline errors"
- Docs: "docs: update README with inline error instructions"
- Test: "test: add tests for handleTransaction"

---

## Quick Commands

- Run app: `go run main.go`
- Run tests: `go test ./...`
- Build: `go build ./...`

---

## Notes

- Keep prompts concise and include the repo path if running against a different workspace.
- When creating new files, include a brief comment header describing purpose.
