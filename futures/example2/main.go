package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Let the main function decide if it's a long running request, then to run in a goroutine
// Returning a channel is still the same
// 		Obvious but send is non-blocking but receive is the blocking side.

func longRunningRequest(r chan<- int32) {
	// Simulate a workload.
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	// Without buffered
	ra, rb := make(chan int32), make(chan int32)
	go longRunningRequest(ra)
	go longRunningRequest(rb)
	fmt.Println(sumSquares(<-ra, <-rb))

	// With buffered channel
	capacity := 2
	buffferedChannel := make(chan int32, capacity)
	for i := 0; i < capacity; i++ {
		go longRunningRequest(buffferedChannel)
	}
	fmt.Println(sumSquares(<-buffferedChannel, <-buffferedChannel))
}
