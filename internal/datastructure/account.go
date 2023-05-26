package datastructure

import (
	"sync"

	"github.com/khatibomar/paytabs-challenge/internal/validator"
)

type Account struct {
	mu      sync.Mutex
	Guid    string
	Name    string
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	v := validator.New()
	err := v.ValidateDeposit(a.Balance, amount)
	if err != nil {
		return err
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	v := validator.New()
	err := v.ValidateWithdrawal(a.Balance, amount)
	if err != nil {
		return err
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance -= amount
	return nil
}

func (a *Account) Transfer(toAccount *Account, amount float64) error {
	v := validator.New()
	err := v.ValidateWithdrawal(a.Balance, amount)
	if err != nil {
		return err
	}
	err = v.ValidateDeposit(toAccount.Balance, amount)
	if err != nil {
		return err
	}

	a.mu.Lock()
	toAccount.mu.Lock()
	defer a.mu.Unlock()
	defer toAccount.mu.Unlock()

	a.Balance -= amount
	toAccount.Balance += amount
	return nil
}

func (a *Account) ValidateAccount() error {
	v := validator.New()
	return v.ValidateBalance(a.Balance)
}
