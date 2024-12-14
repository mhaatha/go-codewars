package main

import (
	"testing"

	"github.com/mhaatha/go-codewars/Pointers-Errors-CS/account"
)

func TestBankAccount(t *testing.T) {
	bankAccount := &account.BankAccount{Owner: "TestUser", Balance: 100}

	t.Run("deposit positive amount", func(t *testing.T) {
		err := bankAccount.Deposit(50)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if bankAccount.GetBalance() != 150 {
			t.Errorf("got %.2f want 150", bankAccount.GetBalance())
		}
	})

	t.Run("deposit negative amount", func(t *testing.T) {
		err := bankAccount.Deposit(-10)
		if err == nil {
			t.Error("expected an error for negative deposit, but got nil")
		}
	})

	t.Run("withdraw valid amount", func(t *testing.T) {
		err := bankAccount.Withdraw(50)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if bankAccount.GetBalance() != 100 {
			t.Errorf("expected balance 100, got %.2f", bankAccount.GetBalance())
		}
	})

	t.Run("withdraw exceeding balance", func(t *testing.T) {
		err := bankAccount.Withdraw(200)
		if err == nil {
			t.Error("expected an error for withdrawal exceeding balance, got nil")
		}
	})
}
