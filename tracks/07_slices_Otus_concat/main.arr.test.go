package main

import "fmt"

func main() {

	v := [5]int{1, 2, 3}
	for i := 0; i < len(v); i++ {
		if v[i] == 0 {
			v[i] = i
		}
	}

	fmt.Println(v)
}
