package main

import (
	"fmt"
	"go-learning/pkg/banking"
)

func main() {
	fmt.Println("=== Valid Withdrawal ===")
	validWithdrawal()

	fmt.Println("\n=== Invalid Withdrawal ===")
	invalidWithdrawal()

	fmt.Println("\n=== Valid Deposit ===")
	validDeposit()

	fmt.Println("\n=== Invalid Deposit ===")
	invalidDeposit()
}

func validWithdrawal() {
	account := &banking.Account{Balance: 100}
	err := banking.ProcessWithdrawal(account, 30)
	if err != nil {
		fmt.Println("Withdrawal error:", err)
		return
	}
	fmt.Printf("New balance: %.2f\n", account.Balance)
}

func invalidWithdrawal() {
	account := &banking.Account{Balance: 50}
	err := banking.ProcessWithdrawal(account, 100)
	if err != nil {
		fmt.Println("Withdrawal error:", err)
		return
	}
	fmt.Printf("New balance: %.2f\n", account.Balance)
}

func validDeposit() {
	account := &banking.Account{Balance: 100}
	err := banking.ProcessDeposit(account, 50)
	if err != nil {
		fmt.Println("Deposit error:", err)
		return
	}
	fmt.Printf("New balance after deposit: %.2f\n", account.Balance)
}

func invalidDeposit() {
	account := &banking.Account{Balance: 100}
	err := banking.ProcessDeposit(account, -100)
	if err != nil {
		fmt.Println("Deposit error:", err)
		return
	}
	fmt.Printf("New balance after deposit: %.2f\n", account.Balance)
}
