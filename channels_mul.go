package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := range 1000 {
			c <- i
		}
		close(c)
	}()

	go func() {
		for v := range c {
			fmt.Println("v=", v, "worker 1")
		}
	}()

	for v := range c {
		fmt.Println("v=", v, "worker 2")
	}

}
