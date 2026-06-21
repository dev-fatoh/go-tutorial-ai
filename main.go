package main

import "fmt"

func main() {
	fmt.Println("Welcome to go bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Withdraw money")
	fmt.Println("3. Deposit money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	accountBalance := 1000.00

	switch choice {
	case 1:
		fmt.Printf("Account balance: %.2f\n", accountBalance)
	case 2:
		var amt float64
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scan(&amt)
		if amt <= 0 {
			fmt.Println("Invalid amount")
		} else if amt > accountBalance {
			fmt.Println("Insufficient funds")
		} else {
			accountBalance -= amt
			fmt.Printf("Withdrawn %.2f. New balance: %.2f\n", amt, accountBalance)
		}
	case 3:
		var amt float64
		fmt.Print("Enter amount to deposit: ")
		fmt.Scan(&amt)
		if amt <= 0 {
			fmt.Println("Invalid amount")
		} else {
			accountBalance += amt
			fmt.Printf("Deposited %.2f. New balance: %.2f\n", amt, accountBalance)
		}
	case 4:
		fmt.Println("Goodbye")
	default:
		fmt.Println("Invalid choice")
	}
}
