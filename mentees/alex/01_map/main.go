package main

import (
	"strconv"
)

func MakeStringDict() map[string]int {
	m := make(map[string]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		m[strconv.Itoa(i)] = i // The map insertions dominate the cost.
	}
	return m
}
