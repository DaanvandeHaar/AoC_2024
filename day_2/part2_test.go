package main

import (
	"testing"
)

func BenchmarkPart2WithFileRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}

func BenchmarkPart2WithFileReadAsync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2Async()
	}
}

func BenchmarkPart2(b *testing.B) {
	reports := ReadFile()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculateSafeReportsPart2(reports)
	}
}
