package main

import (
	"fmt"
	"time"
)

// реализация паттерна генератор
func writer() <-chan int {
	c := make(chan int)

	go func() {
		for i := range 10 {
			c <- i + 1
		}
		close(c) // закончили писать в канал и сразу закрыли его
	}()
	return c
}

// возвращает и принимает канал
func dubbler(ch <-chan int) <-chan int {
	outputChan := make(chan int)

	go func() {
		for i := range ch {
			time.Sleep(500 * time.Millisecond)
			outputChan <- i * 2
		}
		close(outputChan)
	}()
	return outputChan
}

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	reader(dubbler(writer())) // паттерн пайплайн
}
