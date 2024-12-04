package main

func Part1() int {
	grid := ReadFile()

	return GetXMASCount(grid)
}

var XMAS_STR = "XMAS"
var SAMX_STR = "SAMX"

func GetXMASCount(grid []string) int {
	var count int
	gridLen := len(grid)
	for i := range grid {
		rowLen := len(grid[i])
		safeDown := i+3 < gridLen

		for j := range grid[i] {
			safeRight := j+3 < rowLen
			safeLeft := j-3 >= 0

			start := string(grid[i][j])
			if start != "X" && start != "S" {
				continue
			}

			if safeDown {
				down := string([]rune{
					rune(grid[i][j]),
					rune(grid[i+1][j]),
					rune(grid[i+2][j]),
					rune(grid[i+3][j]),
				})
				if down == XMAS_STR || down == SAMX_STR {
					count++
				}
			}

			if safeRight {
				side := grid[i][j : j+4]
				if side == XMAS_STR || side == SAMX_STR {
					count++
				}
			}

			if safeDown && safeRight {
				diagRight := string([]rune{
					rune(grid[i][j]),
					rune(grid[i+1][j+1]),
					rune(grid[i+2][j+2]),
					rune(grid[i+3][j+3]),
				})
				if diagRight == XMAS_STR || diagRight == SAMX_STR {
					count++
				}
			}

			if safeDown && safeLeft {
				diagLeft := string([]rune{
					rune(grid[i][j]),
					rune(grid[i+1][j-1]),
					rune(grid[i+2][j-2]),
					rune(grid[i+3][j-3]),
				})
				if diagLeft == XMAS_STR || diagLeft == SAMX_STR {
					count++
				}
			}
		}
	}

	return count
}
