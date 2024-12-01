package main

import (
	"testing"
)

func BenchmarkPart2WithFileRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}

func BenchmarkPart2(b *testing.B) {
	left, right := ReadFile()
	rightOccurrences := GetOccurrenceMap(right)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculateScore(left, rightOccurrences)
	}
}
