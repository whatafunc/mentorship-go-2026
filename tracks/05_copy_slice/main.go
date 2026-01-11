package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	// a 123
	// b 123 - shared underlying array

	b[0] = 100
	// a 100 2 3 // shared underlying array
	// b 100 2 3 // shared underlying array

	b = append(b, 4)
	// a 100 2 3
	// b 100 2 3 4 // new underlying array

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	// a 100 2 3
	// 100 2 3 4
	modify(a)
	// a 100 2 3
	// b 100 2 3 4
	fmt.Println("after modify a:", a)
}

func modify(s []int) {
	s = append(s, 200) //append creates a new underlying array
	s[1] = 999
}
