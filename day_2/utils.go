package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadFile() [][]int {
	file, err := os.Open("C:\\Users\\goel4\\IdeaProjects\\AoC 2024\\day_2\\input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var numbers [][]int
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = strings.TrimSpace(line)

		parts := strings.Fields(line)
		row := make([]int, len(parts))

		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			row[i] = num
		}
		numbers = append(numbers, row)
	}

	return numbers
}

func ReadFileStream(channel chan []int) {
	file, _ := os.Open("C:\\Users\\goel4\\IdeaProjects\\AoC 2024\\day_2\\input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := bytes.Split(scanner.Bytes(), []byte(" "))
		line := make([]int, len(parts))
		for i, part := range parts {
			p, _ := strconv.Atoi(string(part))
			line[i] = p
		}
		channel <- line
	}

	close(channel)
}
