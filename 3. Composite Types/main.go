package main

import (
	"fmt"
)

func main() {
	//arrays()
	//slices()
	//capacity()
	makeSlices()
}

func arrays() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	var arr2 = [...]string{"Go", "is", "fun"}
	fmt.Println("Array with inferred size:", arr2)
}
func slices() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	slice = append(slice, 6)
	fmt.Println("Appended Slice:", slice)
}
func capacity() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}

func makeSlices() {
	x := make([]int, 0, 10)
	x = append(x, 5, 6, 7, 8)
	fmt.Println(x, len(x), cap(x))
}
