package main

import (
	"AoC_2024/pkg"
	"bufio"
	"bytes"
	"os"
	"strconv"
)

func ReadFile() ([]int, []int) {
	var left, right []int
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := bytes.Split(scanner.Bytes(), []byte("   "))
		left, right = append(left, pkg.Must(strconv.Atoi(string(split[0])))), append(right, pkg.Must(strconv.Atoi(string(split[1]))))
	}

	return left, right
}
