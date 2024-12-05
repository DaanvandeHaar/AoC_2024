package main

import (
	"slices"
)

func Part2() int {
	rules, updates := ReadFile()

	rulesMap := GetRulesMap(rules)

	rejected := GetRejectedUpdates(rulesMap, updates)

	var middles []int
	for _, update := range rejected {
		slices.SortFunc(update,
			func(a, b int) int {
				mustAfter, ok := rulesMap[a]
				if !ok {
					return 0
				}

				if Contains(b, mustAfter) {
					return -1
				}

				return 1
			})
		middles = append(middles, GetMiddle(update))
	}

	var total int
	for _, middle := range middles {
		total += middle
	}

	return total
}

func GetRejectedUpdates(rulesMap map[int][]int, updates [][]int) [][]int {
	var rejected [][]int

	for _, update := range updates {
		var seen []int
		hasError := false
		for _, pageNumber := range update {
			mustAfter, ok := rulesMap[pageNumber]
			if !ok {
				seen = append(seen, pageNumber)
				continue
			}
			if ContainsAny(seen, mustAfter) {
				hasError = true
				break
			}
			seen = append(seen, pageNumber)
		}
		if hasError {
			rejected = append(rejected, update)
		}
	}

	return rejected
}
