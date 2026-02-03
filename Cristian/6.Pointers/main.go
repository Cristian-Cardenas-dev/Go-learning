package main

import (
	"fmt"
)

func main() {
	validateUser()
}

func validateUser() {
	type User struct {
		Name string
		Age  int
	}
	user := User{Name: "Alice", Age: 20}
	fmt.Println("User:", user)
	updateAgeWithPointer(&user.Age)
	// The age in user struct is updated
	fmt.Println("User:", user)
	updateAgeWithoutPointer(user.Age)
	// The age in user struct remains unchanged
	fmt.Println("User:", user)
}
func updateAgeWithPointer(age *int) {
	*age = 25
	fmt.Println("New Age:", *age)
}
func updateAgeWithoutPointer(age int) {
	age = 100
	fmt.Println("New Age:", age)
}
