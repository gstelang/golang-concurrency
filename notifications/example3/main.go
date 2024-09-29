package main

import (
	"log"
	"time"
)

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // block here and wait a notification
	log.Print("Worker#", id, " starts.")
	// Simulate a workload.
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, " job done.")
	// Notify the main goroutine (N-to-1),
	done <- T{}
}

func main() {
	// if you want to turn off date time.
	// log.SetFlags(0)
	ready, done := make(chan T), make(chan T)

	numWorkers := 3

	for i := 0; i < numWorkers; i++ {
		go worker(i, ready, done)
	}

	// Simulate an initialization phase.
	time.Sleep(time.Second * 3 / 2)

	// // 1-to-N notifications. close the channel.
	ready <- T{}
	ready <- T{}
	ready <- T{}
	// OR
	// close(ready)

	// Being N-to-1 notified. Sync.waitgroup is the common practice.
	<-done
	<-done
	<-done
}
