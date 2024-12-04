package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func Part1Async() {
	inputChan := make(chan string)
	workerCount := 6

	go func() {
		ReadFileStream(inputChan)
	}()

	CalculateMultiply(inputChan, workerCount)
}

func CalculateMultiply(inputChan chan string, workerCount int) int {
	done := make(chan struct{})
	var mut sync.Mutex
	var total int

	worker := func() {
		for line := range inputChan {
			pairs := parseString(line)
			for _, pair := range pairs {
				mut.Lock()
				total += pair[0] * pair[1]
				mut.Unlock()
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

	return total
}

func parseString(str string) [][]int {
	var pairs [][]int
	var scanMode bool
	var scanned []uint8

	for i := 0; i < len(str); i++ {
		char := str[i]
		if scanMode {
			switch {
			case char >= '0' && char <= '9':
				scanned = append(scanned, char)
			case char == ',':
				scanned = append(scanned, char)
			case char == ')':
				pair, err := parseMultiply(scanned)
				if err == nil {
					pairs = append(pairs, pair)
				}
				scanMode = false
			default:
				scanMode = false
			}
		} else if i+3 < len(str) {
			if string(str[i:i+4]) == "mul(" {
				scanMode = true
				i += 3
				scanned = []uint8{}
			}
		} else {
			break
		}
	}

	return pairs
}

func parseMultiply(scanned []uint8) ([]int, error) {
	parts := strings.Split(string(scanned), ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("failed to parse")
	}
	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return []int{left, right}, nil
}
