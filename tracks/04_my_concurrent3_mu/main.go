package main

import (
	"fmt"
	"sync"
)

type myData struct {
	counter int
}

func main() {
	mySlovaric := sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 1_000; i++ {
		mySlovaric.Store(i, &myData{counter: 0})
	}
	for i := 0; i < 1_000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val, ok := mySlovaric.Load(i)
			if ok {
				myM := val.(*myData)
				myM.counter++
				fmt.Printf("Worker %d, current = %d\n", i, myM.counter)
			}
		}(i)
	}
	wg.Wait()

	mySlovaric.Range(func(key, value any) bool {
		fmt.Printf("res cnt = %d\n", value.(*myData).counter)
		return true
	})
	//fmt.Printf("myM= %v\n", &myM)

}
