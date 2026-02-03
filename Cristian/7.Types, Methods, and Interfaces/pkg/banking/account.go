package banking

import "errors"

type Money float64

type Account struct {
	Balance Money
}

func (a *Account) Withdraw(amount Money) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	if a.Balance < amount {
		return errors.New("insufficient balance")
	}

	a.Balance -= amount
	return nil
}

func (a *Account) Deposit(amount Money) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	a.Balance += amount
	return nil
}
