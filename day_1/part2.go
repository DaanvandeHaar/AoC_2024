package main

func Part2() int {
	left, right := ReadFile()

	return CalculateScore(left, GetOccurrenceMap(right))
}

func CalculateScore(left []int, rightOccurrences map[int]int) int {
	var tot int
	for _, number := range left {
		tot += number * rightOccurrences[number]
	}

	return tot
}

func GetOccurrenceMap(right []int) map[int]int {
	occurrenceMap := make(map[int]int)
	for _, number := range right {
		occurrenceMap[number] += 1
	}

	return occurrenceMap
}
