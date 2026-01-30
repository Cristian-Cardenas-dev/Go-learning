package main

import (
	"fmt"
)

func main() {
	//arrays()
	//slices()
	//capacity()
	//makeSlices()
	makeMaps()
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
	a := []int{1, 2, 3, 4}
	y := a[:2]
	z := a[1:]
	z[0] = 100
	fmt.Println("a:", a)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
func makeMaps() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println("Map:", m)
	delete(m, "one")
	fmt.Println("Map after deletion:", m)
	mint := map[string]int{
		"a": 0,
	}

	v := m["a"]
	v2 := mint["b"]
	fmt.Println(v, v2)
	if val, ok := mint["b"]; ok {
		fmt.Println("Key 'b' found with value:", val)
	} else {
		fmt.Println("Key 'b' not found")
	}
}
