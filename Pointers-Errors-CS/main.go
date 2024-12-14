package main

import (
	"fmt"

	"github.com/mhaatha/go-codewars/Pointers-Errors-CS/account"
)

func main() {
	// Create a new bank account
	bankAccount := &account.BankAccount{Owner: "Alice", Balance: 0}

	// Perform some operations
	if err := bankAccount.Deposit(100); err != nil {
		fmt.Println("Deposit error:", err)
	}
	if err := bankAccount.Withdraw(50); err != nil {
		fmt.Println("Withdraw error:", err)
	}
	if err := bankAccount.Withdraw(70); err != nil {
		fmt.Println("Withdraw error:", err)
	}

	fmt.Printf("Final balance: %.2f\n", bankAccount.GetBalance())
}
