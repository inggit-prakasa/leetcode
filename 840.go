package main

import (
	"fmt"
	"log"
)

func numMagicSquaresInside(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	total := 0
	sum := 15
	for i := 0; i < row-2; i++ {
		for j := 0; j < col-2; j++ {
			// Check if 5 is not center, return 0 if not
			if grid[i+1][j+1] != 5 {
				continue
			}

			// check all 9 value
			if checkVal(grid, i, j) {
				fmt.Println("WRONG VAL")
				continue
			}

			// row 1
			row1 := grid[i][j] + grid[i][j+1] + grid[i][j+2]
			if sum != row1 {
				log.Println("ROW 1")
				continue
			}

			// row 2
			row2 := grid[i+1][j] + grid[i+1][j+1] + grid[i+1][j+2]
			if sum != row2 {
				log.Println("ROW 2")
				continue
			}

			// row 3
			row3 := grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if sum != row3 {
				log.Println("ROW 3")
				continue
			}

			// col 1
			col1 := grid[i][j] + grid[i+1][j] + grid[i+2][j]
			if sum != col1 {
				log.Println("COL 1")
				continue
			}

			// col 2
			col2 := grid[i][j+1] + grid[i+1][j+1] + grid[i+2][j+1]
			if sum != col2 {
				log.Println("COL 2")
				continue
			}

			// col 3
			col3 := grid[i][j+2] + grid[i+1][j+2] + grid[i+2][j+2]
			if sum != col3 {
				log.Println("COL 3")
				continue
			}

			// diagonal 1
			dia1 := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
			if sum != dia1 {
				log.Println("DIA 1")
				continue
			}

			// diagonal 2
			dia2 := grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j]
			if sum != dia2 {
				log.Println("DIA 2")
				continue
			}

			total++
		}
	}

	return total
}

func checkVal(grid [][]int, i, j int) bool {
	seen := [10]bool{}

	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			val := grid[k+i][l+j]

			if val < 1 || val > 9 {
				return true
			}

			if seen[val] {
				return true
			}

			seen[val] = true
		}
	}

	return false
}
