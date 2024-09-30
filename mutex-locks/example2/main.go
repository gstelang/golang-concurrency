package main

import "fmt"

func main() {
	// THE CAPACITY MUST BE ONE
	mutex := make(chan struct{}, 1)
	mutex <- struct{}{} // this line is needed.

	counter := 0
	increase := func() {
		<-mutex // lock. This will be blocked until the one on line 14 is executed.
		counter++
		mutex <- struct{}{} // unlock
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
