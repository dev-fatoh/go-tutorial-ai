// DOM references used throughout the frontend script.
const balanceEl = document.getElementById("balance");
const form = document.getElementById("transaction-form");
const messageEl = document.getElementById("message");

// setFieldError shows or clears an inline error message for a field.
// It also toggles ARIA attributes and an error CSS class for accessibility.
function setFieldError(fieldId, text) {
  const errEl = document.getElementById(`${fieldId}-error`);
  const inputEl = document.getElementById(fieldId);
  if (!errEl || !inputEl) return;
  errEl.textContent = text || "";
  if (text) {
    inputEl.classList.add("input-error");
    inputEl.setAttribute("aria-invalid", "true");
  } else {
    inputEl.classList.remove("input-error");
    inputEl.removeAttribute("aria-invalid");
  }
}

// clearFieldErrors removes inline errors from all known fields.
function clearFieldErrors() {
  setFieldError("amount", "");
  setFieldError("action", "");
}

// loadBalance fetches the current balance from the API and updates the UI.
async function loadBalance() {
  try {
    const response = await fetch("/api/balance");
    const data = await response.json();
    if (!response.ok) {
      throw new Error(data.error || "Could not load balance");
    }

    balanceEl.textContent = `$${data.balance.toFixed(2)}`;
  } catch (error) {
    showMessage(error.message, "error");
  }
}

// showMessage displays a global status message (success or error).
function showMessage(text, type) {
  messageEl.textContent = text;
  messageEl.className = `message ${type}`;
}

// Form submission: validate, call the API, and show errors inline or global.
form.addEventListener("submit", async (event) => {
  event.preventDefault();

  const action = document.getElementById("action").value;
  const amount = parseFloat(document.getElementById("amount").value);
  clearFieldErrors();

  if (Number.isNaN(amount) || amount <= 0) {
    // Client-side validation for positive numbers.
    setFieldError("amount", "Please enter a valid amount.");
    return;
  }

  try {
    const response = await fetch("/api/transaction", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ action, amount }),
    });

    const data = await response.json();
    if (!response.ok) {
      // If server returned field-specific errors, display them inline
      if (data.errors) {
        Object.keys(data.errors).forEach((field) => {
          setFieldError(field, data.errors[field]);
        });
      }
      throw new Error(data.error || "Transaction failed");
    }

    showMessage(data.message, "success");
    document.getElementById("amount").value = "";
    await loadBalance();
  } catch (error) {
    showMessage(error.message, "error");
  }
});

loadBalance();
