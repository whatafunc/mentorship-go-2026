package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter = 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < 5_000_000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
			fmt.Printf("Worker %d\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Printf("cnt = %d\n", counter)
}
