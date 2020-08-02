package practice

import (
	"fmt"
	"testing"
)

func TestNQueens(t *testing.T) {
	fmt.Println(solveNQueens(8))
}

func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	if n == 0 {
		return ans
	}

	// init
	track := make([][]string, 0)
	for i := 0; i < n; i++ {
		temp := make([]string, 0)
		for j := 0; j < n; j++ {
			temp = append(temp, ".")
		}
		track = append(track, temp)
	}
	backTrackNQueens(&ans, track, 0)
	return ans
}

func backTrackNQueens(ans *[][]string, track [][]string, row int) {
	if row == len(track) {
		temp := make([]string, 0)
		for _, v := range track {
			str := ""
			for _, v2 := range v {
				str += v2
			}
			temp = append(temp, str)
		}
		*ans = append(*ans, temp)
		return
	}

	col := track[row]
	for i := 0; i < len(col); i++ {
		if !isValid(track, row, i) {
			continue
		}
		track[row][i] = "Q"
		backTrackNQueens(ans, track, row+1)
		track[row][i] = "."
	}
}

func isValid(track [][]string, row int, col int) bool {
	// 检查列是否有冲突
	for i := 0; i < len(track); i++ {
		if track[i][col] == "Q" {
			return false
		}
	}

	// 检查左上方是否有冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if track[i][j] == "Q" {
			return false
		}
	}

	// 检查右上方是否有冲突
	for i, j := row-1, col+1; i >= 0 && j < len(track); i, j = i-1, j+1 {
		if track[i][j] == "Q" {
			return false
		}
	}
	return true
}
