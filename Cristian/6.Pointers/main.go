package main

import (
	"errors"
	"fmt"
)

func main() {
	validateUser()

}

func validateUser() (bool, error) {
	type User struct {
		Name string
		Age  int
	}
	user := User{Name: "Alice"}
	fmt.Println("User:", user, "Age address:", &user.Age)
	if _, err := validateAgeUser(&user.Age); err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	return true, nil
}
func validateAgeUser(age *int) (bool, error) {
	if age == nil {
		return false, errors.New("age is invalid")
	}
	if *age < 18 {
		return false, errors.New("user is underage")
	}
	return true, nil
}


