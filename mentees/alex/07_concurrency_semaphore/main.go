// https://vkvideo.ru/video-145052891_456250444?t=2s
// semaphore example as to protect full workers from being used at the same time,
// and to wait for all workers to finish before exiting the program
package main

import (
	"fmt"
	"time"
)

const NumWorkers = 3

func main() {
	tasks := []string{
		"task1", "task2", "task3", "task4",
		"task5", "task6", "task7", "task8",
	}

	sem := make(chan struct{}, NumWorkers)
	done := make(chan struct{}, len(tasks))

	for _, v := range tasks {
		go func(task string) {
			sem <- struct{}{}
			defer func() { <-sem }()

			fmt.Printf("Worker %s started\n", task)
			time.Sleep(3 * time.Second)
			fmt.Printf("Worker %s done\n", task)

			done <- struct{}{}
		}(v)
	}

	for range tasks {
		<-done
	}
}
