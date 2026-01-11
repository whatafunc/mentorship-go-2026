package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {

		defer wg.Done()
		fmt.Println("Hello, 世界")
	}()
	wg.Wait()
}
