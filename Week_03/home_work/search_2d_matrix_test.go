package home_work

import (
	"fmt"
	"testing"
)

func TestSearChMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50}}
	target := 3
	fmt.Println(searchMatrix3(matrix, target))
}

// 暴力求解
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// 变成数组，然后二分 ** 不需要一维转换为二维😀
func searchMatrix2(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}

	matrixAry := make([]int, 0)
	for i := 0; i < row; i++ {
		matrixAry = append(matrixAry, matrix[i]...)
	}

	// 二分
	l := 0
	r := len(matrixAry) - 1
	for l <= r {
		mid := (r + l) / 2
		if matrixAry[mid] == target {
			return true
		}

		if matrixAry[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}

// 可以用转换的方式来访问
func searchMatrix3(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}

	col := len(matrix[0])
	l := 0
	r := col*row - 1
	for l <= r {
		mid := (r + l) / 2
		val := matrix[mid/col][mid%col]
		if target == val {
			return true
		} else {
			if val < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
