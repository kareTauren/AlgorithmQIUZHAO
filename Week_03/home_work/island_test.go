package home_work

import (
	"fmt"
	"testing"
)

func TestNumIsLands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '1', '0'}}

	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}

	col := len(grid[0])
	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}

func dfs(grad [][]byte, row, col int) {
	if !isValid(grad, row, col) {
		return
	}

	if grad[row][col] == '0' {
		return
	}

	grad[row][col] = '0'
	dfs(grad, row-1, col)
	dfs(grad, row+1, col)
	dfs(grad, row, col-1)
	dfs(grad, row, col+1)
}

func isValid(grad [][]byte, row, col int) bool {
	return row >= 0 && row < len(grad) && col >= 0 && col < len(grad[0])
}
