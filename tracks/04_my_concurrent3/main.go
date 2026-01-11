package main

import (
	"fmt"
	"sync"
)

type myData struct {
	counter int
	mu      sync.Mutex
}

func main() {
	mySlovaric := make(map[int]*myData)
	wg := sync.WaitGroup{}
	for i := 0; i < 1_000; i++ {
		mySlovaric[i] = &myData{}
	}
	for i := 0; i < 1_000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myM := mySlovaric[i]
			myM.mu.Lock()
			myM.counter++
			myM.mu.Unlock()
			fmt.Printf("Worker %d, current = %d\n", i, myM.counter)
		}(i)
	}
	wg.Wait()

	for _, myM := range mySlovaric {
		fmt.Printf("res cnt = %d\n", myM.counter)
	}
	//fmt.Printf("myM= %v\n", &myM)

}
