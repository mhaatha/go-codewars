package account

import "errors"

// Deposit adds money to the account
func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	b.Balance += amount
	return nil
}

// Withdraw removes money from the account if sufficient funds are available
func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if amount > b.Balance {
		return errors.New("insufficient funds")
	}
	b.Balance -= amount
	return nil
}

// GetBalance returns the current balance of the account
func (b *BankAccount) GetBalance() float64 {
	return b.Balance
}
