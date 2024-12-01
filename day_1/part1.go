package main

import (
	"slices"
)

func Part1() int {
	left, right := ReadFile()
	slices.Sort(right)
	slices.Sort(left)

	return CalculateDistance(left, right)
}

func CalculateDistance(left []int, right []int) int {
	var tot int
	for i := 0; i < len(left); i++ {
		r, l := left[i], right[i]
		if r < l {
			tot += l - r
		} else {
			tot += r - l
		}
	}

	return tot
}
