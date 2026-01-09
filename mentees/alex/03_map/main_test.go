package main

import (
	"testing"
)

var sink map[int]int

func BenchmarkMapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = MakeIntDict()

	}

	/*
		BenchmarkMapInsert-8   	      15	  89.436.956 ns/op	40.236.382 B/op	      20 allocs/op

	*/
}
