package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter = 0
	wg := sync.WaitGroup{}
	mu := &sync.RWMutex{}
	for i := 0; i < 7_000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
			fmt.Printf("Worker %d working\n", i)
		}(i)
	}

	for i := 0; i < 7000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.RLock()
			defer mu.RUnlock()
			fmt.Printf("cnt = %d\n", counter)
		}()
	}

	wg.Wait()
	println("ok, fyn")
}
