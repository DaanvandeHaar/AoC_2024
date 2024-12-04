package main

import (
	"bufio"
	"os"
)

func ReadFileStream(channel chan string) {
	file, _ := os.Open("C:\\Users\\goel4\\IdeaProjects\\AoC 2024\\day_3\\input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		channel <- scanner.Text()
	}

	close(channel)
}

func ReadFile() []string {
	var lines []string
	file, _ := os.Open("C:\\Users\\goel4\\IdeaProjects\\AoC 2024\\day_3\\input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
