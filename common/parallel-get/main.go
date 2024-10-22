package main

import (
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

// Write to the result channel
func get(url string, ch chan result) {

	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{
			url:     url,
			err:     err,
			latency: time.Since(start).Round(time.Millisecond),
		}
	} else {
		ch <- result{
			url:     url,
			err:     nil,
			latency: time.Since(start).Round(time.Millisecond),
		}
		resp.Body.Close()
	}
}

func main() {

	urls := []string{"https://www.google.com", "https://www.apple.com", "https://www.microsoft.com"}

	ch := make(chan result, len(urls))
	for _, url := range urls {
		go get(url, ch)
	}

	// Slightly ugly but alternate solution without using waitgroup
	// for select required labeled break to break out of for loop
	count := 0
outerLoop:
	for {
		select {
		case val, ok := <-ch:
			if ok {
				count++
				fmt.Println(val.url + ":" + val.latency.String())
				if count == len(urls) {
					// We've processed all the results, and we can safely break
					// use labeled break or return to exit.
					break outerLoop
					//  return
				}
			} else {
				break outerLoop
			}
		default:
		}
	}

	fmt.Println("All processing done")
}
