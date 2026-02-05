package typeErrors

import "fmt"

type DailyLimitExceededError struct {
	Limit  float64
	Amount float64
}

func (e *DailyLimitExceededError) Error() string {
	return fmt.Sprintf(
		"daily limit exceeded: limit=%.2f amount=%.2f",
		e.Limit,
		e.Amount,
	)
}
