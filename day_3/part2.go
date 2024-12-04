package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Part2())
}

func Part2() int {
	lines := ReadFile()

	return ParseLines(lines)
}

func ParseLines(lines []string) int {
	var total int
	enabled := true
	for _, line := range lines {
		rowTotal, e := parseLine(line, enabled)
		total += rowTotal
		enabled = e
	}

	return total
}

func parseLine(str string, enabled bool) (int, bool) {
	var total int
	var scanned []uint8

	var scanMode bool

	for i := 0; i < len(str); i++ {
		char := str[i]
		if scanMode {
			switch {
			case char >= '0' && char <= '9':
				scanned = append(scanned, char)
			case char == ',':
				scanned = append(scanned, char)
			case char == ')':
				mul, err := parseMul(scanned)
				if err == nil {
					total += mul
				}
				scanMode = false
			default:
				scanMode = false
			}
		} else {
			if enabled {
				if i+3 < len(str) {
					if string(str[i:i+4]) == "mul(" {
						scanMode = true
						i += 3
						scanned = []uint8{}
					}
				}
				if i+6 < len(str) {
					if string(str[i:i+7]) == "don't()" {
						scanMode = false
						enabled = false
						i += 3
					}
				}
			} else {
				if i+3 < len(str) {
					if string(str[i:i+4]) == "do()" {
						enabled = true
						i += 3
					}
				}
			}
		}
	}

	return total, enabled
}

func parseMul(scanned []uint8) (int, error) {
	parts := strings.Split(string(scanned), ",")
	if len(parts) != 2 {
		return 0, fmt.Errorf("failed to parse")
	}
	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	return left * right, nil
}
