package main

import (
	"testing"
)

func BenchmarkHttpDI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		diForBench(func(up UserProvider) {})
	}
}
