package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// age := 16
	// valid, err := validateUser(age)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Valid user:", valid)
	// }
	deferExampleOpenFile()
}

func validateUser(age int) (bool, error) {
	if age < 18 {
		return false, errors.New("user is underage")
	}
	return true, nil
}
func deferExampleOpenFile() {
	file, err := os.Open("test.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully")

}
func validateTransfer(amount float64, balance float64) (bool, error) {
	if amount <= 0 {
		return false, errors.New("the amount must be greater than zero")
	}
	if amount > balance {
		return false, errors.New("the amount exceeds the available balance")
	}
	return true, nil
}
