package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Response duration varies per call

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// Sleep 1s/2s/3s.
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

func main() {
	startTime := time.Now()
	// c must be a buffered channel.
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		go source(c)
	}

	// Only the first response will be used.
	rnd := <-c

	// Following will block till all is available.
	// Only the first response will be used.
	// a, b, c1, d, e := <-c, <-c, <-c, <-c, <-c
	// fmt.Println(a + b + c1 + d + e)

	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)

}
