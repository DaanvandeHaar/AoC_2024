package main

import (
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
		left, right = append(left, Must(strconv.Atoi(string(split[0])))), append(right, Must(strconv.Atoi(string(split[1]))))
	}

	return left, right
}

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}
