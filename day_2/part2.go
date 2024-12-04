package main

import "sync"

func isReportSafe(report []int) bool {
	isIncreasing := report[0] < report[1]

	for j := 1; j < len(report); j++ {
		lastSeen, current := report[j-1], report[j]

		if isIncreasing && current < lastSeen {
			return false
		}
		if !isIncreasing && current > lastSeen {
			return false
		}

		if Abs(lastSeen-current) > 3 || lastSeen == current {
			return false
		}
	}

	return true
}

func CalculateSafeReportsPart2(reports [][]int) int {
	var safeCount int
	for _, report := range reports {
		if isReportSafe(report) {
			safeCount++
			continue
		}

		for k := 0; k < len(report); k++ {
			reportWithDampening := make([]int, 0, len(report)-1)
			reportWithDampening = append(reportWithDampening, report[:k]...)
			reportWithDampening = append(reportWithDampening, report[k+1:]...)
			if isReportSafe(reportWithDampening) {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func CalculateSafeReportsPart2Concurrent(inputChan chan []int, workerCount int) int {
	var safeCount int
	var mu sync.Mutex
	done := make(chan struct{})

	worker := func() {
		for report := range inputChan {
			if isReportSafe(report) {
				mu.Lock()
				safeCount++
				mu.Unlock()
				continue
			}

			for k := 0; k < len(report); k++ {
				reportWithDampening := make([]int, 0, len(report)-1)
				reportWithDampening = append(reportWithDampening, report[:k]...)
				reportWithDampening = append(reportWithDampening, report[k+1:]...)
				if isReportSafe(reportWithDampening) {
					mu.Lock()
					safeCount++
					mu.Unlock()
					break
				}
			}
		}
		done <- struct{}{}
	}

	for i := 0; i < workerCount; i++ {
		go worker()
	}

	for i := 0; i < workerCount; i++ {
		<-done
	}

	return safeCount
}

func Part2() int {
	reports := ReadFile()
	return CalculateSafeReportsPart2(reports)
}

func Part2Async() int {
	inputChan := make(chan []int)
	workerCount := 16

	go func() {
		ReadFileStream(inputChan)
	}()

	return CalculateSafeReportsPart2Concurrent(inputChan, workerCount)
}
