package main

import (
	"fmt"
	"go_dev_tour/toy"
)

func main() {
	toy := toy.New("constructor", 11.2)
	fmt.Printf("Toy: %v\n", toy)
	fmt.Println(toy.Name)
	fmt.Println(toy.OnHand(false))
}
