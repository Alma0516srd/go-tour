package main

import "time"

func predict_toWork() {
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()
	select {
	case <-ch:
	case <-time.After(3 * time.Second):

	}
}

func main() {
	predict_toWork()
}

func randomTimeWork() {
	time.Sleep(1 * time.Hour)
}
