package home_work

import (
	"fmt"
	"sort"
	"testing"
)

func TestMinTopK(t *testing.T) {
	arr := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4, 8}
	k := 8
	// fmt.Println(getLeastNumbers(arr, k))
	fmt.Println(getLeastNumbers2(arr, k))
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= 1 {
		return arr
	}

	ans := make([]int, 0)
	tempAns := make([][]int, 10000)
	// for i := 0; i < 10000; i++ {
	// 	tempAns[i] = []int{-1}
	// }
	for i := 0; i < len(arr); i++ {
		tempAns[arr[i]] = append(tempAns[arr[i]], arr[i])
	}
	fmt.Println(tempAns)
	for i := 0; i < len(tempAns) && len(ans) < k; i++ {
		if tempAns[i] == nil {
			continue
		}
		if len(ans)+len(tempAns[i]) > k {
			ans = append(ans, tempAns[i][0:k-len(ans)]...)
		} else {
			ans = append(ans, tempAns[i]...)
		}
	}
	return ans
}

// 排序，取前 k 个
func getLeastNumbers2(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
