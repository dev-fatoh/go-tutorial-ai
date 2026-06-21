package main

import "fmt"

func main() {
	// Starting balance for the account.
	accountBalance := 1000.00

	// Keep showing the menu until the user chooses to exit.
	for {
		// Display the bank menu options.
		fmt.Println("\nWelcome to go bank")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Withdraw money")
		fmt.Println("3. Deposit money")
		fmt.Println("4. Exit")

		// Read the user's menu choice.
		var choice int
		fmt.Print("Enter choice: ")
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid input")
			return
		}

		// Handle the selected banking action.
		switch choice {
		case 1:
			// Show the current account balance.
			fmt.Printf("Account balance: %.2f\n", accountBalance)
		case 2:
			// Withdraw money from the account.
			var amt float64
			fmt.Print("Enter amount to withdraw: ")
			if _, err := fmt.Scan(&amt); err != nil {
				fmt.Println("Invalid input")
				return
			}
			if amt <= 0 {
				fmt.Println("Invalid amount")
			} else if amt > accountBalance {
				fmt.Println("Insufficient funds")
			} else {
				accountBalance -= amt
				fmt.Printf("Withdrawn %.2f. New balance: %.2f\n", amt, accountBalance)
			}
		case 3:
			// Deposit money into the account.
			var amt float64
			fmt.Print("Enter amount to deposit: ")
			if _, err := fmt.Scan(&amt); err != nil {
				fmt.Println("Invalid input")
				return
			}
			if amt <= 0 {
				fmt.Println("Invalid amount")
			} else {
				accountBalance += amt
				fmt.Printf("Deposited %.2f. New balance: %.2f\n", amt, accountBalance)
			}
		case 4:
			// Exit the program.
			fmt.Println("Goodbye")
			return
		default:
			// Handle unsupported menu choices.
			fmt.Println("Invalid choice")
		}
	}
}
