package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int32
	wg := sync.WaitGroup{}
	//mu := sync.Mutex{}
	for i := 0; i < 7_000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//mu.Lock()
			counter.Add(1)
			//mu.Unlock()
			fmt.Printf("Worker %d\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Printf("cnt = %d\n", counter.Load())
}
