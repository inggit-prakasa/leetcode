package main

import "fmt"

func numberOfPaths(grid [][]int, k int) int {
	const modAll = 1e9 + 7
	rows := len(grid)
	cols := len(grid[0])

	dp := make([][][]int, rows)

	for i := range dp {
		dp[i] = make([][]int, cols)
		for j := 0; j < cols; j++ {
			dp[i][j] = make([]int, k)
		}
	}

	dp[0][0][grid[0][0]%k] = 1

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for l := 0; l < k; l++ {

				fmt.Println(i, j,dp[i][j])
				if i > 0 {
					newMod := (l + grid[i][j]) % k
					dp[i][j][newMod] = (dp[i][j][newMod] + dp[i-1][j][l]) % modAll
				}
				if j > 0 {
					newMod := (l + grid[i][j]) % k
					dp[i][j][newMod] = (dp[i][j][newMod] + dp[i][j-1][l]) % modAll
				}
			}
		}
	}

	fmt.Println(dp)

	return dp[rows-1][cols-1][0]
}
