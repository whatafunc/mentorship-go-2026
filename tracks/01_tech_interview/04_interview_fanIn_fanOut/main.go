package main

import (
	"fmt"
	"sync"
	"time"
)

// fetch simulates a blocking HTTP fetch call.
func fetch(url string, workerId int) (int, error) {
	fmt.Println("Fetching (and 'working') per URL:", url, "by worker", workerId)

	time.Sleep(100 * time.Millisecond) // Simulate network delay

	// In real code, you would use http.Get or similar
	return 1, nil
}

func main() {
	urls := []string{
		"http://example.com",
		"http://example.org",
		"http://example.net",
		"http://example.edu",
		"http://example.io",
		"http://example.co.uk",
		"http://example.travel",
	}
	inputChan := make(chan string, len(urls)) // Buffered channel to hold URLs
	resultsChan := make(chan int, len(urls))  // Buffered channel to hold results
	for _, url := range urls {
		inputChan <- url
	}
	close(inputChan)               // Close the channel after sending all URLs - asap
	workers := make([]struct{}, 3) // 3 workers
	wg := sync.WaitGroup{}
	for i := range workers {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for v := range inputChan { // fanOut
				result, err := fetch(v, id) // fetch is a blocking call
				if err != nil {
					fmt.Printf("Error fetching %s: %v\n", v, err)
				} else {
					resultsChan <- result // fanIn
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultsChan) // Close results channel after all workers are done
	}()

	// Collect results
	fmt.Println("Fan-in results:")
	for res := range resultsChan {
		fmt.Println("Result received:", res)
	}
	fmt.Println("Finished")
}
