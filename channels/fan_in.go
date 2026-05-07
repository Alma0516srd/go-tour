package main

import (
	"fmt"
	"sync"
)

func worker1(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		out <- n * n
	}
}

func main() {
	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	// 3 workers (fan-out)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker1(input, output, &wg)
	}

	// closer
	go func() {
		wg.Wait()
		close(output)
	}()

	// sender
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	// fan-in
	for res := range output {
		fmt.Println(res)
	}

}
