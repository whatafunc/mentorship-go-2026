package main

import (
	"testing"
)

var sink map[string]int

func BenchmarkMapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = MakeStringDict()

	}

	/*
		goos: darwin
		goarch: amd64
		cpu: Intel(R) Core(TM) i7-4750HQ CPU @ 2.00GHz
		BenchmarkMapInsert-8   	       6	 272.498.913 ns/op	65.653.118 B/op	  999903 allocs/op
	*/
}
