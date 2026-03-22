package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	const RES_THRESHOLD_SIZE = 100
	const RANDOM_SEED = 5000

	c := make(chan int)
	shutdown := make(chan struct{})
	var poolSize = runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(poolSize)

	// отдельный генератор случайных чисел для каждой горутины
	for i := 0; i < poolSize; i++ {
		go func(id int) {
			rnd := rand.New(rand.NewSource(int64(RANDOM_SEED + id)))
			for {
				rndVal := rnd.Intn(RANDOM_SEED)

				select {
				case c <- rndVal:
					fmt.Println("send", rndVal)
				case <-shutdown:
					fmt.Println("shutdown")
					wg.Done()
					return
				}
			}
		}(i)
	}

	var res []int
	for val := range c {
		if val%2 == 0 {
			fmt.Println("Skip value")
			continue
		}
		fmt.Println("Store value to slice")
		res = append(res, val)

		if len(res) == RES_THRESHOLD_SIZE {
			break
		}
	}

	close(shutdown)
	wg.Wait()
	fmt.Println(res)
}
