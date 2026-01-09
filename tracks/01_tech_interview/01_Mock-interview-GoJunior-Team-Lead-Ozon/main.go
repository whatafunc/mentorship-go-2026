package main

import (
	"fmt"
	"sync"
	"time"
)

// Account struct for slice demo
type Account struct{ ID int }

// --- 1. Slice Demo ---
func sliceDemo() {
	fmt.Println("--- Slice Demo ---")
	s1 := make([]Account, 0, 2)
	s1 = append(s1, Account{ID: 1})
	s2 := append(s1, Account{ID: 2})
	s2[0].ID = 42 // Mutating s2 affects s1 because they share the underlying array
	fmt.Printf("s1: %v (len=%d, cap=%d)\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))
	fmt.Println()
}

// --- 2. String Demo ---
func stringDemo() {
	fmt.Println("--- String Demo ---")
	s := "Hello, 世界"
	// s[0] = 'h' // Compile error: strings are immutable

	// Correct way to "modify" a string
	sBytes := []byte(s)
	sBytes[0] = 'h'
	s = string(sBytes)
	fmt.Println("Modified string:", s)

	// Iterating over a string (yields runes)
	for i, r := range s {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}
	fmt.Println()
}

// --- 3. Map Demo ---
func mapDemo() {
	fmt.Println("--- Map Demo ---")
	// Maps are not safe for concurrent use without a mutex
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	// Deleting a key
	delete(m, "a")

	// Checking for existence
	if val, ok := m["b"]; ok {
		fmt.Println("Key 'b' exists, value:", val)
	}
	if _, ok := m["a"]; !ok {
		fmt.Println("Key 'a' does not exist")
	}
	fmt.Println()
}

// --- 4. Interface Demo ---
type CustomError struct{}

func (e *CustomError) Error() string {
	return "custom error"
}

func getError() error {
	var err *CustomError // err is a nil pointer of type *CustomError
	return err           // The returned interface contains a type, so it's not nil
}

func interfaceDemo() {
	fmt.Println("--- Interface Demo ---")
	err := getError()
	// An interface is only nil if both its type and value are nil.
	// Here, the type is *CustomError and the value is nil.
	fmt.Printf("err is nil? %v (Type: %T, Value: %v)\n", err == nil, err, err)
	fmt.Println()
}

// --- 5. Defer and Closure Demo ---
type X struct {
	Value int
}

func deferDemo() {
	fmt.Println("--- Defer and Closure Demo ---")
	x := &X{Value: 123}

	// defer evaluates arguments immediately. It prints the value of x.Value at the time of the defer statement.
	defer fmt.Println("Defer 1 (immediate eval):", x.Value)

	// This closure captures 'x' by reference. It prints the value of x.Value when the function exits.
	defer func() {
		fmt.Println("Defer 2 (closure):", x.Value)
	}()

	x.Value = 456
	fmt.Println("Function end")
	fmt.Println()
}

// --- 6. Safe Counter (Concurrency) ---
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func safeCounterDemo() {
	fmt.Println("")
	fmt.Println("--- Safe Counter Demo ---")
	counter := SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		// wg.Go(func() { //go 1.25+
		// 	counter.Inc()
		// 	//.Println("cur counter value:", counter.Value())
		// })
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
			//fmt.Println("cur counter value:", counter.Value())
		}()

	}
	wg.Wait()
	fmt.Println("Final counter value:", counter.Value())
	fmt.Println()
}

// --- 7. Fan-In Pattern (Merge Channels) ---
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close `out` once all the output goroutines are done.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fanInDemo() {
	fmt.Println("--- Fan-In Pattern Demo ---")
	// Create a few input channels
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	// Start goroutines to send data to the channels
	go func() {
		defer close(c1)
		for i := 0; i < 3; i++ {
			c1 <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		defer close(c2)
		for i := 10; i < 13; i++ {
			c2 <- i
			time.Sleep(time.Millisecond * 150)
		}
	}()
	go func() {
		defer close(c3)
		for i := 20; i < 23; i++ {
			c3 <- i
			time.Sleep(time.Millisecond * 50)
		}
	}()

	// fmt.Println("performing merge using mergeChan function:")
	// out := mergeChan(c1, c2, c3)
	// for n := range out {
	// 	fmt.Printf("Received: %d\n", n)
	// }

	fmt.Println("performing merge using merge function:")
	// Merge the channels and read from the merged channel
	merged := merge(c1, c2, c3)
	for n := range merged {
		fmt.Printf("Received: %d\n", n)
	}
	fmt.Println("Done receiving from merged channel.")
	fmt.Println()
}

func mergeChan(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, ch := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// sliceDemo()
	// stringDemo()
	// mapDemo()
	// interfaceDemo()
	// deferDemo()
	// safeCounterDemo()
	fanInDemo()
}
