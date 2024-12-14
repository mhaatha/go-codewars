package account

// Account defines the interface for managing a bank account
type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

// BankAccount represents a simple bank account
type BankAccount struct {
	Owner   string
	Balance float64
}
