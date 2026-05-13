package main

import (
	"fmt"
)

func joinChans(chans ...<-chan int32) <-chan int32 {
	//out := make(chan int32, 100) //if we run without goroutine, we need to buffer the channel, otherwise it will block
	out := make(chan int32) // if we run with goroutine, we can use unbuffered channel, because we will read from it in the same time as we write to it
	go func() {             // need goroutine to avoid deadlock, because we need to read from chans and write to out at the same time
		defer close(out)
		for _, ch := range chans {
			for v := range ch {
				out <- v
			}
		}
	}()
	return out
}

func main() {
	a := make(chan int32, 10)
	b := make(chan int32, 6)
	for i := 0; i < 10; i++ {
		a <- int32(i)
		if i < 6 {
			b <- int32(i)
		}
	}
	close(a)
	close(b)
	for i := range joinChans(a, b) {
		fmt.Println(i)
	}

}
