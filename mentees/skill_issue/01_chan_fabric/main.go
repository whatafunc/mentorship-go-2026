package main

func producer() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			println("sent = ", i)
		}
		close(ch)
	}()
	return ch
}

func consumer(ch chan int) chan int {
	chRes := make(chan int, 10)
	for val := range ch {
		chRes <- val * 2
		println("consumed val: ", val)
	}
	close(chRes)
	return chRes
}

func printerRes(ch chan int) {
	for val := range ch {
		println("printing val: ", val)
	}
}

func main() {
	println("Hello, World!")
	ch := make(chan int, 10)

	for val := range producer() {
		ch <- val
		println("got val: ", val)
	}
	close(ch)

	println("launch consumer")

	printerRes(consumer(ch))
	// for val := range ch {
	// 	println("consumed val: ", val)
	// }

	println("FYN!")
}
