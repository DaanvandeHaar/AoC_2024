package main

import "testing"

func BenchmarkPart1WithFileRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}
