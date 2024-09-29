package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	// Creates a byte slice of 32 MB
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // can be buffered or not

	// The sorting goroutine
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		// Notify sorting is done.
		done <- struct{}{}
	}()

	// do some other things ...

	<-done // waiting here for notification

	// As a note, you don't need to do explicit conversion because byte is an alias for uint8
	// When you pass a byte to fmt.Println, it automatically converts it to its integer representation (0-255) for printing.
	// This means that you don’t need to explicitly convert it to an integer type; Go handles this conversion internally.
	fmt.Println(values[0], values[len(values)-1])

}
