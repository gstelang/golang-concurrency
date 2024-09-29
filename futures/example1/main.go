package main

import (
	"fmt"
	"math/rand"
	"time"
)

// longRunningRequest shouldn't block and does so by returning a channel.

// 1. Run the long running job in a goroutine...duh!
// 1.1 Wrtie to the channel once result is available.
// 2. Return a receive only channel.

// longRunningRequest will take about ~3 seconds instead of 3+3

func longRunningRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		// Simulate a workload.
		fmt.Println("Working hard to generate a number")
		time.Sleep(time.Second * 3)
		// generates a number between 0 to 99
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a * a
}

func main() {
	a := longRunningRequest()
	b := longRunningRequest()
	// As a side note, the following will print the memory address of a
	fmt.Println(a)
	fmt.Println(sumSquares(<-a, <-b))
}
