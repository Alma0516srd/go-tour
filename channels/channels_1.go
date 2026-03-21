package main

import (
	"fmt"
	"sync"
)

func main() {
	const wgCount = 2
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(wgCount)

	go func() {
		doSomething(c)
		wg.Done() // If we don't decrement wg counter - we get deadlock
	}()

	go func() {
		doSomething(c)
		wg.Done()
	}()

	c <- 1
	wg.Wait()
}

func doSomething(channel chan int) {
	for {

		value, status := <-channel
		if !status {

			fmt.Println("Goroutine is down")
			return
		}

		fmt.Println("Goroutine has value ", value)

		if value == 10 {
			close(channel)
			fmt.Println("Channel was closed")
			return
		}

		// send to channel
		value = value + 1
		channel <- value
	}
}
