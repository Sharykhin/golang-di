package main

import "testing"

func BenchmarkDI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DI(handlerT)
	}
}
