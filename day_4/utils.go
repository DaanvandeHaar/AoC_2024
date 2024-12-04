package main

import (
	"bufio"
	"os"
)

func ReadFile() []string {
	var lines []string

	file, _ := os.Open("C:\\Users\\goel4\\IdeaProjects\\AoC 2024\\day_4\\input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
