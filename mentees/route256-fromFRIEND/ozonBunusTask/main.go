package main

import (
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan int, 10)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i + 1
		}
		//out <- 0
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Merge2Channels(func(x int) int { return x }, ch1, ch2, out, 10)
	}()

	//go func() {
	//defer wg.Done()
	//close(out)
	//close(ch2)

	//}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range out {
			println("Out:", v)
		}
	}()
	wg.Wait()
	println("FYN")
}

func Merge2Channels(
	f func(int) int,
	in1 <-chan int,
	in2 <-chan int,
	out chan<- int,
	n int) {

	for i := 0; i < n; i++ {
		out <- f(<-in1) + f(<-in2)
		println("Merging worked #", i)
	}

	println("_____merging___completed_____")
	defer close(out)

}
