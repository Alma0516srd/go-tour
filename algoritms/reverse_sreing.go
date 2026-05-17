package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverse(s)
	fmt.Println(string(s)) // "olleh"
}

func reverse(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp

		left++
		right--
	}
}
