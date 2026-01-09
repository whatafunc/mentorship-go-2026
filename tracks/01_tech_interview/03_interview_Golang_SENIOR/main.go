package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var nums = []int{5, 2, 4, 1, 3}

func main() {
	fmt.Println("--- StartWorkers basic ---")
	StartWorkers(nums)
	fmt.Println("--- StartWorkers Diagnosted version ---")
	StartWorkersDiagnosted(nums)
	fmt.Println("--- Deterministic version ---")
	StartWorkersDeterministic(nums)
	fmt.Println("--- Deterministic version B ---")
	StartWorkersDeterministicB(nums)
}

func StartWorkers(nums []int) {
	ch := make(chan int, len(nums))
	wg := sync.WaitGroup{}

	// start workers first so they are ready to receive values and output results
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for n := range ch { // it is Go runtime that schedules which goroutine gets value
				// from channel
				fmt.Println(id, n*n) // print worker id and squared number
			}
		}(i)
	}

	// send numbers and then close channel to signal workers
	for _, n := range nums { // отправляем числа в канал aka feed the channel
		ch <- n
	}
	close(ch)

	wg.Wait()
}

// 2 25
// 1 1
// 2 16
// 0 4
// 1 9

func StartWorkersDiagnosted(nums []int) {
	numWorkers := 3
	ch := make(chan int, len(nums))
	wg := sync.WaitGroup{}
	var counts [3]int32

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for n := range ch {
				fmt.Println(id, n*n)
				atomic.AddInt32(&counts[id], 1)
			}
		}(i)
	}

	for _, n := range nums {
		ch <- n // отправляем числа в канал aka feed the channel
	}
	close(ch)
	wg.Wait()

	// print counts for each worker as a Quick diagnostic
	for i := 0; i < numWorkers; i++ {
		fmt.Printf("worker %d processed %d items\n", i, counts[i])
	}
}

// --- Debug version output ---
// 0 4
// 0 16
// 0 9
// 2 25
// 1 1
// worker 0 processed 3 items
// worker 1 processed 1 items
// worker 2 processed 1 items

func StartWorkersDeterministic(nums []int) {
	numWorkers := 3
	workers := make([]chan int, numWorkers)
	wg := sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		workers[i] = make(chan int)
		wg.Add(1)
		go func(id int, ch <-chan int) {
			defer wg.Done()
			for n := range ch {
				fmt.Println(id, n*n)
			}
		}(i, workers[i])
	}

	// dispatch deterministically by index
	for idx, n := range nums {
		workers[idx%numWorkers] <- n // sends the number n to the corresponding worker's channel.
	}
	for i := 0; i < numWorkers; i++ { // closing each worker's channel. This is usually normal practice done when:
		// No more tasks will be sent to the worker.
		// You want the worker goroutine to finish after processing all pending tasks.
		close(workers[i])
	}
	wg.Wait()
}

// --- Deterministic version ---
// 0 25
// 0 1
// 1 4
// 1 9
// 2 16

func StartWorkersDeterministicB(nums []int) {
	type result struct {
		idx      int
		workerID int
		value    int
	}
	results := make([]result, len(nums))
	resCh := make(chan result)
	numWorkers := 3
	wg := sync.WaitGroup{}

	// worker pool reading from shared channel (or dispatcher)
	jobs := make(chan struct {
		idx int
		n   int
	}, len(nums))

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				resCh <- result{idx: job.idx, workerID: id, value: job.n * job.n}
			}
		}(i)
	}

	// send jobs
	for idx, n := range nums {
		jobs <- struct {
			idx int
			n   int
		}{idx, n}
	}
	close(jobs)

	// collect results concurrently
	go func() {
		wg.Wait()
		close(resCh)
	}()

	for r := range resCh {
		results[r.idx] = r
	}

	// print in original order
	for _, r := range results {
		fmt.Println(r.workerID, r.value)
	}
}
