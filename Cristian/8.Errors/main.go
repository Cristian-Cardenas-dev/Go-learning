package main

import (
	"errors"
	"fmt"
	"go-errors/pkg/typeErrors"
)

var (
	ErrInvalidAmount  = errors.New("invalid amount")
	ErrAccountBlocked = errors.New("account is blocked")
)

func main() {
	fmt.Println("=== Invalid amount ===")
	if err := CanPay(-100, 50, false); err != nil {
		HandlePaymentError(err)
	}
	fmt.Println("\n=== Daily limit exceeded ===")
	if err := CanPay(200, 50, false); err != nil {
		HandlePaymentError(err)
	}
	fmt.Println("\n=== Account blocked ===")
	if err := CanPay(20, 50, true); err != nil {
		HandlePaymentError(err)
	}
	fmt.Println("\n=== Valid payment ===")
	if err := CanPay(20, 50, false); err != nil {
		HandlePaymentError(err)
	} else {
		fmt.Println("Payment processed successfully")
	}

}

func CanPay(amount, dailyLimit float64, accountBlocked bool) error {
	if accountBlocked {
		return fmt.Errorf("can pay failed: %w", ErrAccountBlocked)
	}
	if amount <= 0 {
		return fmt.Errorf("can pay failed: %w", ErrInvalidAmount)
	}
	if amount > dailyLimit {
		return &typeErrors.DailyLimitExceededError{
			Limit:  dailyLimit,
			Amount: amount,
		}
	}
	return nil
}

func HandlePaymentError(err error) {
	var limitErr *typeErrors.DailyLimitExceededError
	switch {
	case errors.Is(err, ErrAccountBlocked):
		fmt.Println("Account is blocked. Please contact support.")
	case errors.Is(err, ErrInvalidAmount):
		fmt.Println("The amount provided is invalid.")
	case errors.As(err, &limitErr):
		fmt.Printf(
			"You exceeded your daily limit (limit %.2f, tried %.2f)\n",
			limitErr.Limit,
			limitErr.Amount,
		)
	default:
		fmt.Println("An unknown error occurred:", err)
	}
}
