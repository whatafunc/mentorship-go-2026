package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	n := m // n = {a:1, b:2} - shared underlying data

	// m map[a:1 b:2]
	// n map[a:1 b:2] - shared underlying data

	n["a"] = 100

	// m map[a:100 b:2] - shared underlying data
	// n map[a:100 b:2] - shared underlying data

	n["c"] = 3
	// m map[a:100 b:2 c:3] - shared underlying data
	// n map[a:100 b:2 c:3] - shared underlying data

	fmt.Println("m:", m)
	fmt.Println("n:", n)

	modifyMap(m)

	fmt.Println("after modify m:", m)
	// m map[a:100 b:2 c:3] - shared underlying data
	fmt.Println("n:", n) // 200 2 3
}

func modifyMap(mp map[string]int) {
	mp["b"] = 200
	mp = map[string]int{ // creates a new map
		"x": 9,
	}
	mp["y"] = 10
}
