package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(Part2())
}

func Part2() int {
	grid := ReadFile()

	return GetXMASCount(grid)
}

var XMAS_BYTES = []byte("XMAS")
var SAMX_BYTES = []byte("SAMX")

func GetXMASCount(grid [][]byte) int {
	var count int
	gridLen := len(grid)
	for i := range grid {
		rowLen := len(grid[i])
		safeDown := i+3 < gridLen

		for j := range grid[i] {
			safeRight := j+3 < rowLen
			safeLeft := j-3 >= 0

			if safeDown {
				down := []byte{grid[i][j], grid[i+1][j], grid[i+2][j], grid[i+3][j]}
				if bytes.Equal(bytes.ToUpper(down), XMAS_BYTES) || bytes.Equal(bytes.ToUpper(down), SAMX_BYTES) {
					count++
				}
			}

			if safeRight {
				side := grid[i][j : j+4]
				if bytes.Equal(bytes.ToUpper(side), XMAS_BYTES) || bytes.Equal(bytes.ToUpper(side), SAMX_BYTES) {
					count++
				}
			}

			if safeDown && safeRight {
				diag := []byte{grid[i][j], grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3]}
				if bytes.Equal(bytes.ToUpper(diag), XMAS_BYTES) || bytes.Equal(bytes.ToUpper(diag), SAMX_BYTES) {
					count++
				}
			}

			if safeDown && safeLeft {
				diag := []byte{grid[i][j], grid[i+1][j-1], grid[i+2][j-2], grid[i+3][j-3]}
				if bytes.Equal(bytes.ToUpper(diag), XMAS_BYTES) || bytes.Equal(bytes.ToUpper(diag), SAMX_BYTES) {
					count++
				}
			}
		}
	}

	return count
}
