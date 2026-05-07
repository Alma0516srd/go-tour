package main

import "fmt"

func worker_pool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		results <- j * 2
	}
}

func main() {

	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for w := 1; w <= 3; w++ {
		go worker_pool(w, jobs, results)
	}

	// заполняем канал
	for _, t := range tasks {
		jobs <- t
	}
	close(jobs)

	for a := 1; a <= len(tasks); a++ {
		fmt.Println(<-results)
	}

}
