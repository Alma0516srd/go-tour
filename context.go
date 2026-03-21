package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Add imports.

type Key int

// Declare a new type named `key` that is based on an int.

const userIPKey Key = 0

// Declare a constant named `userIPKey` of type `key` set to
// the value of 0.

type myUser struct {
	Name  string
	Email string
}

// Declare a struct type named `User` with two `string` based
// fields named `Name` and `Email`.

func main() {
	routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// routes sets the routes for the web service.
func routes() {
	http.HandleFunc("/user", findUser)
}

// Implement the findUser function to leverage the context for
// both timeouts and state.
func findUser(rw http.ResponseWriter, r *http.Request) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancelFunc()

	// Create a context that timeouts in fifty milliseconds.

	// Defer the call to cancel.

	// Save the `r.RemoteAddr` value in the context using `userIPKey`
	// as the key. This call returns a new context so replace the
	// current `ctx` value with this new one. The original context is
	// the parent context for this new child context.
	ctx = context.WithValue(ctx, userIPKey, r.RemoteAddr)
	channel := make(chan *myUser, 1)

	// Create a channel with a buffer size of 1 that works with
	// pointers of type `User`

	// Use this goroutine to make the database call. Use the channel
	// to get the user back.
	go func() {
		if ip, ok := ctx.Value(userIPKey).(string); ok {
			log.Println("Connected to external db", ip)
		}
		// Get the `r.RemoteAddr` value from the context and log
		// the value you get back.

		// Call the `readDatabase` function provided below and
		// send the returned `User` pointer on the channel.
		channel <- readDatabase()
		log.Println("goroutine is terminating")
		// Log that the goroutine is terminating.
	}()

	// Wait for the database call to finish or the timeout.
	select {
	case u := <-channel:

		// Add a case to wait on the channel for the `User` pointer.

		// Call the `sendResponse` function provided below to
		// send the `User` to the caller. Use `http.StatusOK`
		// as the status code.

		sendResponse(rw, u, http.StatusOK)
		log.Println("Response ", rw)

		// Log we sent the response with a StatusOk

		return

		// Add a case to wait on the `ctx.Done()` channel.

		// Use this struct value for the error response.
		e := struct{ Error string }{ctx.Err().Error()}
		print(e)

		// Call the `sendResponse` function provided below to
		// send the error to the caller. Use `http.StatusRequestTimeout`
		// as the status code.

		// Log we sent the response with a StatusRequestTimeout

		return
	}
}

// readDatabase performs a pretend database call with
// a second of latency.
func readDatabase() *myUser {
	u := myUser{
		Name:  "Bill",
		Email: "bill@ardanlabs.com",
	}

	// Create 100 milliseconds of latency.
	time.Sleep(100 * time.Millisecond)

	return &u
}

// sendResponse marshals the provided value into json and returns
// that back to the caller.
func sendResponse(rw http.ResponseWriter, v interface{}, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(v)
}
