package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile() ([][2]int, [][]int) {
	var rules [][2]int
	var pages [][]int

	content, err := os.ReadFile("C:\\Users\\goel4\\IdeaProjects\\AoC 2024\\day_5\\input.txt")
	if err != nil {
		log.Fatal("could not read file")
	}

	fileContent := string(content)
	parts := strings.Split(fileContent, "\r\n\r\n")
	if len(parts) != 2 {
		log.Fatal("incorrect file content")
	}

	part1 := strings.Split(parts[0], "\r\n")
	for _, line := range part1 {
		ruleSet := strings.Split(line, "|")
		rules = append(rules, [2]int{Must(strconv.Atoi(ruleSet[0])), Must(strconv.Atoi(ruleSet[1]))})
	}

	part2 := strings.Split(parts[1], "\r\n")
	for _, line := range part2 {
		pageSet := strings.Split(line, ",")

		var page []int
		for _, n := range pageSet {
			page = append(page, Must(strconv.Atoi(n)))
		}
		pages = append(pages, page)
	}

	return rules, pages
}

func GetRulesMap(rules [][2]int) map[int][]int {
	rulesMap := make(map[int][]int)
	for _, rule := range rules {
		rulesMap[rule[0]] = append(rulesMap[rule[0]], rule[1])
	}

	return rulesMap
}

func GetAllowedUpdatesMiddle(rulesMap map[int][]int, updates [][]int) []int {
	var middles []int

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
		if !hasError {
			middles = append(middles, GetMiddle(update))
		}
	}

	return middles
}

func GetMiddle(update []int) int {
	return update[len(update)/2]
}

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}

func Contains[T comparable](comparable T, array []T) bool {
	for _, item := range array {
		if item == comparable {
			return true
		}
	}

	return false
}

func ContainsAny[T comparable](set []T, contains []T) bool {
	lookup := make(map[T]struct{})

	for _, item := range set {
		lookup[item] = struct{}{}
	}

	for _, item := range contains {
		if _, exists := lookup[item]; exists {
			return true
		}
	}

	return false
}
