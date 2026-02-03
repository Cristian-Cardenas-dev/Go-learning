package banking

func ProcessWithdrawal(a *Account, amount Money) error {
	return a.Withdraw(amount)
}

func ProcessDeposit(a *Account, amount Money) error {
	return a.Deposit(amount)
}
