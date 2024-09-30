package main

import (
	"context"
	"fmt"
	"time"
)

func AfterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func main() {
	fmt.Println("Hi!")
	<-AfterDuration(time.Second)
	fmt.Println("Hello!")
	<-AfterDuration(time.Second)
	fmt.Println("Bye!")

	ch := time.After(2 * time.Second)
	fmt.Println(<-ch) // Will print the current time after 2 seconds

	timer := time.NewTimer(3 * time.Second)
	<-timer.C                    // Blocks until 3 seconds have passed, and then receives the time
	timer.Stop()                 // Stops the timer if it hasn't fired yet
	timer.Reset(2 * time.Second) // Resets the timer to a new duration

	ticker := time.NewTicker(1 * time.Second)
	count := 0
	for t := range ticker.C {
		fmt.Println(t) // Prints the time every second
		if count == 10 {
			ticker.Stop()
		}
	}

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Called after 2 seconds")
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Timeout reached")
	}

}
