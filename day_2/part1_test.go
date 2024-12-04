package main

import "testing"

func BenchmarkPart1WithFileRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkPart1(b *testing.B) {
	reports := ReadFile()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculateSafeReportsPart1(reports)
	}
}
