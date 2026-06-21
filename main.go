package main

import "fmt"

// main is the entry point for this simple CLI bank application.
// It prints a menu, reads the user's choice, and performs the selected action.
func main() {
	// Display menu options
	fmt.Println("Welcome to go bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Withdraw money")
	fmt.Println("3. Deposit money")
	fmt.Println("4. Exit")

	// Read the user's menu choice
	var choice int
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	// Starting account balance (in a real app this would be persisted/stored)
	accountBalance := 1000.00

	// Handle user selection
	switch choice {
	case 1:
		// Show current balance
		fmt.Printf("Account balance: %.2f\n", accountBalance)
	case 2:
		// Withdraw flow: read amount and validate
		var amt float64
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scan(&amt)
		if amt <= 0 {
			// Negative or zero amounts are invalid
			fmt.Println("Invalid amount")
		} else if amt > accountBalance {
			// Cannot withdraw more than current balance
			fmt.Println("Insufficient funds")
		} else {
			// Subtract amount from balance and report
			accountBalance -= amt
			fmt.Printf("Withdrawn %.2f. New balance: %.2f\n", amt, accountBalance)
		}
	case 3:
		// Deposit flow: read amount and validate
		var amt float64
		fmt.Print("Enter amount to deposit: ")
		fmt.Scan(&amt)
		if amt <= 0 {
			// Negative or zero amounts are invalid
			fmt.Println("Invalid amount")
		} else {
			// Add amount to balance and report
			accountBalance += amt
			fmt.Printf("Deposited %.2f. New balance: %.2f\n", amt, accountBalance)
		}
	case 4:
		// Exit the program
		fmt.Println("Goodbye")
	default:
		// Any other input is an invalid menu choice
		fmt.Println("Invalid choice")
	}
}
