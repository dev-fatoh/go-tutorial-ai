# Go Bank CLI

A simple command-line banking program written in Go.

## Overview

This project demonstrates a basic bank menu where a user can:

- check the current balance,
- deposit money,
- withdraw money, and
- exit the program.

The program uses a loop so the menu keeps appearing until the user chooses to exit.

## Features

- Check account balance
- Deposit money
- Withdraw money
- Exit the program
- Input validation for invalid menu choices and negative amounts

## How to Run

1. Open the project folder.
2. Make sure Go is installed on your system.
3. Run the program with:

   ```bash
   go run main.go
   ```

## Example Flow

When you run the program, you will see a menu like this:

```text
1. Check balance
2. Withdraw money
3. Deposit money
4. Exit
```

If you choose option `3`, the program will ask for the deposit amount. If the amount is valid, the balance is updated and shown on screen.

## Notes

- The balance starts at `1000.00` when the program begins.
- The program currently stores the balance only while the program is running.
- This is a beginner-friendly example for learning Go basics such as loops, conditionals, and user input.
