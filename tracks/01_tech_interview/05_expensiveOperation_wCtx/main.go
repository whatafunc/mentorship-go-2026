package main

import (
	"context"
	"sync"
	"time"
)

func expensiveOperation(ctx context.Context, item string) (string, error) { // fix: ? need to return err??
	// Simulate an expensive operation
	ch := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
	}()
	select {
	case <-ctx.Done():
		return "ctx exited", ctx.Err()
	case <-ch:
		return item + "_processed", nil
	}
}

func ProcessItems(ctx context.Context, items []string) []string {
	results := []string{}
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, item := range items {
		wg.Add(1)
		go func(currentItem string) { // fix2: get input as argument to avoid closure capture issue
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
				processed, err := expensiveOperation(ctx, currentItem)
				if err != nil {
					return
				}

				mu.Lock()                            // fix3:  Added sync.Mutex to protect concurrent access to the results slice
				results = append(results, processed) // fix: add iterator? // fix: check error?
				mu.Unlock()
			}

		}(item) // fix1: pass item as argument to avoid closure capture issue
	}

	wg.Wait()
	return results
}

// Credits to:
// https://www.youtube.com/watch?v=SfMIpBM2Xf8
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	start := time.Now()
	items := []string{"item1", "item2", "item3", "item4", "item5"}
	processedItems := ProcessItems(ctx, items)
	for _, item := range processedItems {
		println(item) //itemX_processed
	}
	println(" FYN , processed items count:", time.Since(start).String())
}
