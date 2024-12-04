package main

func Part2() int {
	grid := ReadFile()
	return GetXMASCountX(grid)
}

var MAS_STR = "MAS"
var SAM_STR = "SAM"

func GetXMASCountX(grid []string) int {
	var count int

	gridLen := len(grid)
	rowLen := len(grid[1])
	for i := 1; i+1 < gridLen; i++ {
		for j := 1; j+1 < rowLen; j++ {
			middle := grid[i][j]

			if middle != 'A' {
				continue
			}

			diagRight := string([]byte{grid[i-1][j-1], middle, grid[i+1][j+1]})
			if !(diagRight == MAS_STR) && !(diagRight == SAM_STR) {
				continue
			}

			diagLeft := string([]byte{grid[i-1][j+1], middle, grid[i+1][j-1]})
			if !(diagLeft == MAS_STR) && !(diagLeft == SAM_STR) {
				continue
			}

			count++
		}
	}

	return count
}
