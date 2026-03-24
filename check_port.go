package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

const (
	PORT_COUNT      = 65536
	HOST            = "127.0.0.1"
	GORUTINES_COUNT = 150
)

type result struct {
	portNumber int
	isOpen     bool
}

func checkPort(portNumber int) bool {
	var addr = HOST + ":" + strconv.Itoa(portNumber)
	connection, err := net.Dial("tcp", addr)
	if err == nil {
		fmt.Println("Connection successful")
		connection.Close()
		return true
	}
	return false
}

func worker(jobs <-chan int, results chan<- result, wg *sync.WaitGroup) {
	defer wg.Done() // после каждого прохода - декремент счетчика
	for port := range jobs {
		isOpen := checkPort(port)
		var res = result{portNumber: port, isOpen: isOpen}
		results <- res // канал только для записи
	}
}

func main() {
	fmt.Println("Инициализация")
	workers := make(chan int, GORUTINES_COUNT) //канал источник данных
	results := make(chan result, GORUTINES_COUNT)
	var wg sync.WaitGroup
	for i := 0; i < GORUTINES_COUNT; i++ {
		wg.Add(1)
		go worker(workers, results, &wg) // запускаем параллельно горутины
	}

	go func() {
		for port := 1; port < PORT_COUNT; port++ {
			workers <- port
		}
		close(workers)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Сканирование портов....")

	for rs := range results {
		fmt.Println(rs)
	}
}
