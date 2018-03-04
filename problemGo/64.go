package main

import (
	"fmt"
)

//Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
//
//Note: You can only move either down or right at any point in time.
//
//Example 1:
//[[1,3,1],
// [1,5,1],
// [4,2,1]]

func minPathSum(grid [][]int) int {
	minn := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = minn(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1][m-1]
}

func main() {
	a := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	res := minPathSum(a)
	fmt.Println(res)
}
