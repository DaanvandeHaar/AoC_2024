package main

import "testing"

func BenchmarkPart1WithFileRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}

func BenchmarkPart1(b *testing.B) {
	left, right := ReadFile()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateDistance(left, right)
	}
}
