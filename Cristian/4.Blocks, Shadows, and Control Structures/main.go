package main

import (
	"fmt"
)

func main() {
	//blockScope()
	// shadowing()
	//ifInGo()
	//the4FormsOfGoFor()
	//breakAndContinueLabels()
	switchInGo()
}

func blockScope() {
	var name string = "hello"
	{
		var name string = "block scope"
		fmt.Println(name)
	}
	fmt.Println(name)
}

func shadowing() {
	var value int = 10
	fmt.Println("Outer value:", value)

	{
		value := 20
		fmt.Println("Inner value:", value)
	}

	fmt.Println("Outer value after inner block:", value)
}

func ifInGo() {
	// Basic if-else statement
	var number int = 5
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}
	// Using short statement in if
	if x := 10; x > 5 {
		fmt.Println(x, "is greater than 5")
	}
}

func the4FormsOfGoFor() {
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Traditional for loop:", i)
	}
	// While-like for loop
	j := 0
	for j < 5 {
		fmt.Println("While-like for loop:", j)
		j++
	}
	// Infinite for loop with break
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println("Infinite loop with break:", k)
		k++
	}
	// For range loop
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

func breakAndContinueLabels() {
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue OuterLoop
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}

func switchInGo() {
	// Basic switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}
	// Switch with fallthrough
	grade := 'B'
	switch grade {
	case 'A':
		fmt.Println("Excellent")
		fallthrough
	case 'B':
		fmt.Println("Well done")
		fallthrough
	case 'C':
		fmt.Println("Good")
	default:
		fmt.Println("Keep trying")
	}
	// Switch with multiple expressions
	fruit := "apple"
	switch fruit {
	case "apple", "banana":
		fmt.Println("This is a common fruit")
	case "kiwi", "mango":
		fmt.Println("This is an exotic fruit")
	default:
		fmt.Println("Unknown fruit")
	}
	// Switch without an expression
	number := 7
	switch {
	case number < 0:
		fmt.Println("Negative number")
	case number == 0:
		fmt.Println("Zero")
	case number > 0:
		fmt.Println("Positive number")
	}
}
