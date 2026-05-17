package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target)) // [0 1]
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range nums {
		complainment := target - v
		if idx, ok := seen[complainment]; ok {
			return []int{idx, i}
		}
		seen[v] = i
	}
	return nil
}
