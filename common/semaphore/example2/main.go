package main

import (
	"fmt"
	"reflect"
	"sync"
)

// Semaphore type that encapsulates the semaphore logic.
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore creates a new Semaphore with a given maximum count.
func NewSemaphore(maxCount int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, maxCount),
	}
}

// Acquire waits to obtain a permit from the semaphore.
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

// Release returns a permit to the semaphore.
func (s *Semaphore) Release() {
	<-s.ch
}

type Executor interface {
	Execute(args ...interface{}) interface{}
}

// Executes any arbitrary function
func (s *Semaphore) Execute(args ...interface{}) interface{} {
	s.Acquire()
	defer s.Release()

	fmt.Println("Executing function....")
	if len(args) == 0 {
		fmt.Println("No function provided")
		return nil
	}

	fn := reflect.ValueOf(args[0])
	if fn.Kind() != reflect.Func {
		fmt.Println("First argument is not a function")
		return nil
	}

	fnArgs := args[1:]
	reflectedArgs := make([]reflect.Value, len(fnArgs))
	for i, arg := range fnArgs {
		reflectedArgs[i] = reflect.ValueOf(arg)
	}

	results := fn.Call(reflectedArgs)
	if len(results) == 0 {
		return nil
	}
	return results[0].Interface()
}

func main() {
	// Define the number of resources and the number of goroutines.
	const numTokens = 3
	const numGoroutines = 9
	semaphore := NewSemaphore(3)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	squareFunc := func(a int) int {
		return a * a
	}

	results := make(chan int)

	// 9 goroutines with 3 executing at any point in time.
	for i := 1; i <= numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			results <- semaphore.Execute(squareFunc, id).(int)
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

}
