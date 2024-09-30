package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

func CountNumbers() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// without this, you will get a deadlock.
		close(ch)
	}()
	return ch
}

func RandomGenerator() <-chan uint64 {
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
				break
			}
			// converts the byte slice into a uint64 using big-endian byte order
			c <- binary.BigEndian.Uint64(rnds)
			time.Sleep(1 * time.Second)
		}
	}()
	return c
}

func main() {
	// this counts upto 10
	ch := <-CountNumbers()
	for val := range ch {
		fmt.Println("Hello : ", val)
	}

	// this runs in an infinite loop
	randomCh := RandomGenerator()
	for val := range randomCh {
		fmt.Println("Random: ", val)
	}

}
