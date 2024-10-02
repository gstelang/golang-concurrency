package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const maxInt = 100000000

var arr = fillArr(maxInt)
var done = make(chan bool, 1)
var result = make(chan string, 1)

func generateRandomNumber() int {
	return rand.Intn(maxInt) + 1
}

func fillArr(maxSize int) []int {
	arr := []int{}
	for i := 0; i < maxSize; i++ {
		arr = append(arr, generateRandomNumber())
	}

	return arr
}

func writeToResult(str string) {
	select {
	case result <- str:
	default:
	}
}

func search(goRoutineId int, target int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine: %d \n", goRoutineId)

	for _, val := range arr {
		select {
		case val, ok := <-done:
			if !ok || val == true {
				fmt.Printf("Goroutine returning %d\n", goRoutineId)
				return
			}
		default:
			if val == target {
				done <- true
				// signal close immediately
				close(done)
				writeToResult("Found it")
				fmt.Printf("Goroutine: %d found it\n", goRoutineId)
				return
			}
		}

	}

	writeToResult("Not Found")
}

func main() {

	numGoRoutines := 5
	target := 989204
	var wg sync.WaitGroup
	// sliceSize := len(arr) / numGoRoutines

	for i := 1; i <= numGoRoutines; i++ {
		// split the array into equal chunks
		// start := id * sliceSize
		// end := start + sliceSize
		// if id == numGoRoutines-1 {
		// // Last goroutine takes any remaining elements
		// 	end = len(arr)
		// }
		// and pass arr[start:end]
		wg.Add(1)
		go search(i, target, &wg)
	}

	wg.Wait()
	fmt.Println(<-result)
	close(result)
}


// run with race 
// go run --race main.go
// you should see the output like
// Goroutine: 1 
// Goroutine: 5 
// Goroutine: 4 
// Goroutine: 2 
// Goroutine: 3 
// Goroutine returning 4
// Goroutine returning 2
// Goroutine returning 3
// Goroutine: 1 found it
// Goroutine returning 5
// Found it
