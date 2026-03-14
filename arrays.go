package main

import "fmt"

func main() {
	var five [5]string

	names := [5]string{"Jack", "Tom", "Artur", "Linda", "Sara"}

	five = names

	for _, v := range five {
		fmt.Println(v, &v)
	}
}
