package main

import "fmt"

func main() {
	slice := make([]int, 0)

	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fruts := []string{"Apple", "Orange", "Banana", "Pear", "Strawberry"}

	for index, val := range fruts {
		fmt.Println(index, val)
	}

	fmt.Println("===================")

	testSlice := fruts[1:3] //take 1 and 2 element, not include 3 [)

	for index, val := range testSlice {
		fmt.Println(index, val)
	}
}
