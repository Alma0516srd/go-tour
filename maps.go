package main

import "fmt"

func main() {
	stringMap := map[string]int{}
	stringMap["Banana"] = 1
	stringMap["Apple"] = 2
	stringMap["Tomato"] = 3
	stringMap["Orange"] = 4
	stringMap["Watermelon"] = 5

	for key, value := range stringMap {
		fmt.Println(key, value)
	}
}
