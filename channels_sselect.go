package main

import (
	"fmt"
	"time"
)

// select - блокирующий оператор
func main() {
	timer := time.NewTimer(1 * time.Millisecond)

	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case v := <-ch1:
		fmt.Println("v= ", v, "from ch1")
	case v := <-ch2:
		fmt.Println("v= ", v, "from ch2")
	case <-timer.C: // повисит секунду и выйдет

	}
}
