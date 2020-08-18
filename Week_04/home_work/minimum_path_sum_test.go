package home_work

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	paths := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	fmt.Println(minPathSum(paths))
}

/**
1. 找子问题,就是走过路径的最小值加上当前值
2.dp[i][j]
3.dp 方程
	a. 当 i > 0, j = 0, 则 dp[i][0] = dp[i-1][0] + grid[i][0]
	b. 当 j > 0, i = 0, 则 dp[0][j] = dp[0][j-1] + grid[0][j]
	c. 当 j > 0, j > 0, 则 dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
*/
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m < 0 {
		return 0
	}
	n := len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	fmt.Println(dp)

	min := func(n, m int) int {
		if n > m {
			return m
		}
		return n
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] += dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}
