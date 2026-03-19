package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for count := 100; count >= 0; count-- {
			fmt.Printf("[A gorutine:%d]\n", count)
		}

		wg.Done() // decrement wautgroup counter
	}()

	go func() {
		for count := 0; count <= 100; count++ {
			fmt.Printf("[B gorutine:%d]\n", count)
		}
		wg.Done() // decrement wautgroup counter
	}()

	wg.Wait() // wait all gorutines here
	fmt.Println("Terminating Program")
}
