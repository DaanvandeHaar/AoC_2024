package main

import "testing"

func BenchmarkPart2WithFileRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
