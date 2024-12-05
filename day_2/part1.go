package main

import "AoC_2024/pkg"

func Part1() int {
	reports := ReadFile()

	return CalculateSafeReportsPart1(reports)
}

func CalculateSafeReportsPart1(reports [][]int) int {
	var safeCount int
	for _, report := range reports {
		isIncreasing := report[0] < report[1]
		isSafe := true

		for j := 1; j < len(report); j++ {
			lastSeen, current := report[j-1], report[j]
			if pkg.Abs(lastSeen-current) > 3 || lastSeen == current {
				isSafe = false
				break
			}

			if isIncreasing && current < lastSeen {
				isSafe = false
				break
			}
			if !isIncreasing && current > lastSeen {
				isSafe = false
				break
			}
		}

		if isSafe {
			safeCount++
		}
	}

	return safeCount
}
