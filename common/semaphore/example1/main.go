package main

import (
	"fmt"
	"sync"
	"time"
)

// In a worker pool pattern, allow only 3 go routines to enter
func main() {
	// Define the number of resources and the number of goroutines.
	const numTokens = 3
	const numGoroutines = 10
	// Create a buffered channel to act as a semaphore, with 3 slots.
	semaphore := make(chan struct{}, numTokens)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 1; i <= numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			semaphore <- struct{}{} // Someone gets a token.

			fmt.Printf("Goroutine %d is acquired the token.\n", id)
			time.Sleep(time.Second) // Simulate heavy work

			<-semaphore // Someone releases the token.
		}(i)
	}

	wg.Wait()
}
