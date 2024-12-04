package main

import (
	"bytes"
)

//func main() {
//	fmt.Println(Part1())
//}

func Part1() int {
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

			start := grid[i][j]
			if start != 'X' && start != 'S' {
				continue
			}

			if safeDown {
				down := []byte{start, grid[i+1][j], grid[i+2][j], grid[i+3][j]}
				if bytes.Equal(down, XMAS_BYTES) || bytes.Equal(down, SAMX_BYTES) {
					count++
				}
			}

			if safeRight {
				side := grid[i][j : j+4]
				if bytes.Equal(side, XMAS_BYTES) || bytes.Equal(side, SAMX_BYTES) {
					count++
				}
			}

			if safeDown && safeRight {
				diagRight := []byte{start, grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3]}
				if bytes.Equal(diagRight, XMAS_BYTES) || bytes.Equal(diagRight, SAMX_BYTES) {
					count++
				}
			}

			if safeDown && safeLeft {
				diagLeft := []byte{start, grid[i+1][j-1], grid[i+2][j-2], grid[i+3][j-3]}
				if bytes.Equal(diagLeft, XMAS_BYTES) || bytes.Equal(diagLeft, SAMX_BYTES) {
					count++
				}
			}
		}
	}

	return count
}
