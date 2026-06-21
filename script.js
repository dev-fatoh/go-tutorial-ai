const balanceEl = document.getElementById("balance");
const form = document.getElementById("transaction-form");
const messageEl = document.getElementById("message");

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

function showMessage(text, type) {
  messageEl.textContent = text;
  messageEl.className = `message ${type}`;
}

form.addEventListener("submit", async (event) => {
  event.preventDefault();

  const action = document.getElementById("action").value;
  const amount = parseFloat(document.getElementById("amount").value);

  if (Number.isNaN(amount) || amount <= 0) {
    showMessage("Please enter a valid amount.", "error");
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
