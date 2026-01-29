package main

import (
	"fmt"
)

func main() {
	//literals()
	theZeroValue()
}

func literals() {
	var intLiteral int = 42
	var floatLiteral float64 = 3.14
	var stringLiteral string = "Go is fun!"
	var boolLiteral bool = true

	fmt.Println("Integer Literal:", intLiteral)
	fmt.Println("Float Literal:", floatLiteral)
	fmt.Println("String Literal:", stringLiteral)
	fmt.Println("Boolean Literal:", boolLiteral)
}
func theZeroValue() {
	var a int
	var b float64
	var c string
	var d bool

	fmt.Printf("Zero value of int: %d\n", a)
	fmt.Printf("Zero value of float64: %f\n", b)
	fmt.Printf("Zero value of string: '%s'\n", c)
	fmt.Printf("Zero value of bool: %t\n", d)
}
