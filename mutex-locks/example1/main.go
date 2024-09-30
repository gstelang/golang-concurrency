package main

import "fmt"

// To create a lock, create a buffered channel of capacity 1
// Now write to it.
// Nobody else can write to it unless that is read. So the block between is the place where you can do the shared variable thing.

// Dont think this is a good pattern but just one of the neat things about using channels.
func main() {
	// THE CAPACITY MUST BE ONE
	mutex := make(chan struct{}, 1)

	counter := 0
	increase := func() {
		mutex <- struct{}{} // Lock because no one else write to it unless the value is read
		counter++
		<-mutex // Release
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)

	<-done
	<-done
	fmt.Println(counter) // 2000
}
