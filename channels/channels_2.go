package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const BUFFER_SIZE = 2000

	intChannel := make(chan int, BUFFER_SIZE)

	for i := 0; i < BUFFER_SIZE; i++ {
		go func() {
			intChannel <- rand.Int()
		}()
	}
	gorutineCount := BUFFER_SIZE
	for gorutineCount > 0 {
		fmt.Println(<-intChannel)
		gorutineCount--
	}
}
