package practice

import (
	"fmt"
	"testing"
)

func TestNumIsLands(t *testing.T) {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '1'},
		[]byte{'0', '0', '0', '1', '0'}}

	// dfs
	// fmt.Println(numIslands(grid))
	// bfs
	fmt.Println(bfsIslands(grid))
}

func numIslands(grid [][]byte) int {
	islands := 0
	r := len(grid)
	c := len(grid[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				islands++
				// dfs_easyUnderstand(grid, i, j)
				dfsIslands(grid, i, j)
			}
		}
	}
	return islands
}

/**
				(-1,c)
(r,-1)  (r, c) 	(r,1)
		  	(1,c)
*/

// DFS容易理解的版本
var directions = [][]int{
	[]int{-1, 0},
	[]int{0, -1},
	[]int{1, 0},
	[]int{0, 1}}

func inArea(grid [][]byte, x int, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func dfs_easyUnderstand(grid [][]byte, r, c int) {
	grid[r][c] = '0'
	for i := 0; i < 4; i++ {
		newX := r + directions[i][0]
		newY := c + directions[i][1]
		if inArea(grid, newX, newY) && grid[newX][newY] == '1' {
			dfs_easyUnderstand(grid, newX, newY)
		}
	}
}

// 方法二
func dfsIslands(grid [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	dfsIslands(grid, i-1, j)
	dfsIslands(grid, i+1, j)
	dfsIslands(grid, i, j-1)
	dfsIslands(grid, i, j+1)
}

func bfsIslands(grid [][]byte) int {
	island := 0
	rows := len(grid)
	if rows == 0 {
		return island
	}
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				island++
				queue := make([]int, 0)
				queue = append(queue, i*cols+j)
				grid[i][j] = '0'

				for len(queue) > 0 {
					val := queue[0]
					queue = queue[1:]
					x := val / cols
					y := val % cols
					for k := 0; k < 4; k++ {
						newX := x + directions[k][0]
						newY := y + directions[k][1]
						if inArea(grid, newX, newY) && grid[newX][newY] == '1' {
							queue = append(queue, newX*cols+newY)
							grid[newX][newY] = '0'
						}
					}
				}
			}
		}
	}
	return island
}
