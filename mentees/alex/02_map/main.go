package main

func MakeIntDict() map[int]int {
	m := make(map[int]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		m[i<<12] = i // The map insertions dominate the cost.
	}
	return m
}
