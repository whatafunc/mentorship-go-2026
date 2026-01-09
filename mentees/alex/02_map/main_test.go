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
		goos: darwin
		goarch: amd64
		cpu: Intel(R) Core(TM) i7-4750HQ CPU @ 2.00GHz
		BenchmarkMapInsert-8   	      12	  83.719.662 ns/op	40.236.384 B/op	      20 allocs/op


	*/
}
