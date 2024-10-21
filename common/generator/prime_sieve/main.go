package main

import (
	"fmt"
)

func generate(ch chan int) {
	for i := 2; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func filter(in chan int, out chan int, prime int) {
	for num := range in {
		if num%prime != 0 {
			out <- num
		}
	}
	close(out)
}

func main() {
	input := make(chan int)
	go generate(input)

	for {
		// First it is read by this line.
		// Rest is read by the filter
		// filtered set getting less and less until there's nothing.
		prime, ok := <-input // 2, 3, 5
		if !ok {
			break
		}
		fmt.Println(prime) // 2, 3, 5

		out := make(chan int)
		go filter(input, out, prime)
		// filter out everything from initial seq {3,4...100} that is divisible by 2
		// filter out everything from {5, 7, 9, .... 100 } that is divisible by 3
		// filter out everything from {7,11..... 100 } that is divisible by 5
		input = out
	}
}
