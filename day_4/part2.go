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
	for i := range grid {
		safeDown := i+2 < gridLen

		for j := 0; j < rowLen; j++ {
			safeRight := j+2 < rowLen

			if safeRight && safeDown {
				diagRight := string([]rune{
					rune(grid[i][j]),
					rune(grid[i+1][j+1]),
					rune(grid[i+2][j+2]),
				})
				if !(diagRight == MAS_STR) && !(diagRight == SAM_STR) {
					continue
				}

				diagLeft := string([]rune{
					rune(grid[i+2][j]),
					rune(grid[i+1][j+1]),
					rune(grid[i][j+2]),
				})
				if !(diagLeft == MAS_STR) && !(diagLeft == SAM_STR) {
					continue
				}

				count++
			}
		}
	}

	return count
}
