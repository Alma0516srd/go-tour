package main

import "fmt"

type restrictionInterface interface {
	string | int
}

func copyfy[T restrictionInterface](source []T) []T {
	dest := make([]T, len(source))
	copy(dest, source)
	return dest
}

func main() {
	phones := []string{"Samsung", "Motorola", "Philips"}
	dest := copyfy(phones)
	fmt.Println(phones)
	fmt.Println(dest)

	nums := []int{1, 2, 3}
	numsDest := copyfy(nums)
	fmt.Println(nums)
	fmt.Println(numsDest)
}
