package main

import (
	"fmt"
	"net/http"
)

var nextID = make(chan int)

func handler(w http.ResponseWriter, r *http.Request) {
	// Handle the request
	fmt.Fprintf(w, "Hello, your ID is %d\n", <-nextID)
}
func counter() {
	for i := 0; ; i++ {
		nextID <- i
	}
}

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	// Handle the request
	fmt.Fprintf(w, "Hello, your ID is %d\n", <-ch)
}
func nextCounter(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {

	// version 1
	go counter()
	// Register the handler function for the root path
	http.HandleFunc("/", handler)

	// version 2: OO version of the program
	// Based off, https://www.youtube.com/watch?v=zJd7Dvg3XCk&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=23
	// Declare a channel of type nextCh
	var ch nextCh = make(chan int)
	go nextCounter(ch)
	// use the method as the handler
	http.HandleFunc("/next", ch.handler)

	// Start the server on port 8080
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
