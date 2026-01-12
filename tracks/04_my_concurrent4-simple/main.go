package main

import (
	"fmt"
	"sync"
	"time"
)

func veryLongComputation() int8 {
	time.Sleep(2 * time.Second)
	return 1
}

func main() {
	locker := sync.Mutex{}
	results := make(chan int8, 100)
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello, 世界")
	}()

	var totelTimes int8
	totelTimes = 0
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- veryLongComputation()
			locker.Lock()
			totelTimes += <-results
			locker.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Done in", time.Since(start))
	fmt.Println("Total results:", totelTimes)
}

// Hello, 世界
// Done in 2s
// Total results: 100
// Note: The output may vary based on the execution time of goroutines.
// The program demonstrates concurrent execution with a mutex to safely accumulate results.
// The output will show the total results after all computations are done.
