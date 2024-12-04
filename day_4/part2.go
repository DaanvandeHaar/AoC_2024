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
	return GetXMASCountX(grid)
}

var MAS_BYTES = []byte("MAS")
var SAM_BYTES = []byte("SAM")

func GetXMASCountX(grid [][]byte) int {
	var count int
	gridLen := len(grid)
	for i, lines := range grid {
		rowLen := len(grid[i])
		safeDown := i+2 < gridLen

		for j := range lines {
			safeRight := j+2 < rowLen

			if grid[i][j] != 'S' && grid[i][j] != 'M' {
				continue
			}

			if safeRight && safeDown {
				diagRight := []byte{grid[i][j], grid[i+1][j+1], grid[i+2][j+2]}
				if !bytes.Equal(diagRight, MAS_BYTES) && !bytes.Equal(diagRight, SAM_BYTES) {
					continue
				}

				diagLeft := []byte{grid[i+2][j], grid[i+1][j+1], grid[i][j+2]}
				if !bytes.Equal(diagLeft, MAS_BYTES) && !bytes.Equal(diagLeft, SAM_BYTES) {
					continue
				}

				count++
			}
		}
	}

	return count
}
