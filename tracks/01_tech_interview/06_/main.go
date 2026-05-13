package main

import (
	"fmt"
	"math/rand"
	"time"
)

// реализовать функцию processParallel
// прокинуть контекст

func processData(v int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return v * 2
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	start := time.Now()
	processParallel(in, out, 5)

	for v := range out {
		fmt.Println("Output in Main: v =", v)
	}
	fmt.Println("main duration:", time.Since(start))
}

func processParallel(in, out chan int, numWorkers int) {
	// Реализация должна быть здесь
	for i := 0; i < numWorkers; i++ {
		go func() {
			for v := range in {
				result := processData(v)
				out <- result
			}
		}()
	}

	go func() {
		// Ждем завершения всех воркеров
		time.Sleep(60 * time.Second) // Заглушка для ожидания завершения
		close(out)
	}()
}
