package home_work

import (
	"fmt"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	matrix := [][]byte{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0}}
	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	if m < 1 || n < 1 || matrix == nil {
		return 0
	}
	min := func(x, z, y int) int {
		if x > y {
			x = y
		}
		if x > z {
			return z
		}
		return x
	}
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	maxSide := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			fmt.Println(matrix[i][j] == 1)

			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				}
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}
