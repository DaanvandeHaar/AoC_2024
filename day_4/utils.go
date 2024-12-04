package main

import (
	"bufio"
	"os"
)

func ReadFile() [][]byte {
	var runes [][]byte
	file, _ := os.Open("C:\\Users\\goel4\\IdeaProjects\\AoC 2024\\day_4\\input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runes = append(runes, scanner.Bytes())
	}

	return runes
}
