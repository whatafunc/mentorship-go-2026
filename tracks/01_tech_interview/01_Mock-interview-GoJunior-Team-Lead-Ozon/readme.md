# Go Interview Summary & Code Examples

This document summarizes the core Go topics discussed in the mock interview video and provides runnable code examples for each concept, which can be found in `main.go`.

**To run the examples, execute:** `go run main.go`

---

## Code Examples (`main.go`)

The `main.go` file contains practical demonstrations of the following concepts:

1.  **`sliceDemo()`**: Shows how appending to a slice can lead to shared underlying arrays and unexpected mutations.
.
2.  **`stringDemo()`**: Illustrates the immutability of strings and the correct way to "modify" them using byte slices. It also shows how `range` iterates over runes.
.
3.  **`mapDemo()`**: Demonstrates basic map operations, including deleting keys and checking for key existence.
4.  **`interfaceDemo()`**: Highlights the common pitfall where an interface holding a `nil` pointer is not itself `nil`.
.
5.  **`deferDemo()`**: Contrasts the immediate evaluation of `defer` arguments with the use of a closure to capture a variable's final state.
.
6.  **`safeCounterDemo()`**: Implements a thread-safe counter using a `sync.Mutex` to prevent race conditions.
.
7.  **`fanInDemo()`**: Provides a complete implementation of the fan-in concurrency pattern, merging multiple channels into a single output channel using goroutines and a `sync.WaitGroup`.

---

## Conceptual Summary

### 1. Strings and Immutability
- **Key Concept:** Strings in Go are immutable byte slices. You cannot change a specific character without creating a new string.
- **Iteration:** Using `range` over a string iterates by runes (Unicode code points), not raw bytes.
- **See:** `stringDemo()` in `main.go`.

### 2. Slices and Memory Management
- **Key Concept:** Slices are dynamic wrappers around arrays. When a slice's capacity is exceeded, Go allocates a new, larger underlying array.
- **Shared Memory:** Multiple slices can point to the same underlying array, leading to non-obvious side effects when one slice is modified.
- **See:** `sliceDemo()` in `main.go`.

### 3. Interfaces and Nil Values
- **Common Pitfall:** An interface is only `nil` if both its underlying type and value are `nil`.
- **Example:** A function returning `(*MyStruct)(nil)` as an `error` results in an interface that is not `nil` because it still holds the type information `*MyStruct`.
- **See:** `interfaceDemo()` in `main.go`.

### 4. Defer, Closures, and Pointers
- **Key Insight:** `defer` evaluates its function arguments at the time the `defer` statement is executed, not when the function call is executed.
- **Solution:** To capture the value of a variable at the time the surrounding function exits, use a closure.
- **See:** `deferDemo()` in `main.go`.

### 5. Concurrency: Mutexes and Maps
- **Key Concept:** Go maps are **not** safe for concurrent read/write operations.
- **Solution:** Use a `sync.Mutex` to guard access to a map that will be used by multiple goroutines. The `defer mu.Unlock()` pattern is best practice.
- **See:** `safeCounterDemo()` in `main.go`.

### 6. Concurrency: Fan-In Pattern
- **Task:** Merge multiple input channels into a single output channel.
- **Solution:** Use a `sync.WaitGroup` to orchestrate the closing of the output channel after all input channels have been fully drained.
- **See:** `fanInDemo()` in `main.go`.

---

**Original Video:**
[Mock-собеседование по Go (Junior) от Team Lead из Ozon](https://www.youtube.com/watch?v=5-rENjTvYeU)