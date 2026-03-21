package main

import (
	"math/rand"
	"sync"
)

func main() {
	const goroutines = 100
	const random_seed = 1000
	intChannel := make(chan int, goroutines)

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		// one rnd number per one gorutine
		go func() {
			defer wg.Done()
			randValue := rand.Intn(random_seed)
			if randValue%2 == 0 {
				return
			}
			intChannel <- randValue
		}()
	}
}
