package home_work

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	// fmt.Println("combine ", combine(4, 2))
	fmt.Println(alaall([]int{1, 2, 3}))
}

// func combine(n int, k int) [][]int {
// 	if n <= 1 {
// 		return [][]int{}
// 	}
//
// 	ans := make([][]int, 0)
// 	for i := 1; i <= n; i++ {
// 		for j := i + 1; j <= n; j++ {
// 			ans = append(ans, []int{i, j})
// 		}
// 	}
// 	return ans
// }

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	if n <= 0 || k <= 0 {
		return ans
	}
	tempAns := make([]int, 0)
	findCombine(n, k, 1, tempAns, &ans)
	return ans
}

func findCombine(n, k int, begin int, tempAns []int, ans *[][]int) {
	if len(tempAns) == k {
		*ans = append(*ans, append([]int{}, tempAns...))
		return
	}

	for i := begin; i <= n; i++ {
		tempAns = append(tempAns, i)
		findCombine(n, k, i+1, tempAns, ans)
		tempAns = tempAns[:len(tempAns)-1]
	}
}

func alaall(nums []int) [][]int {
	ans := make([][]int, 0)
	tempAns := make([]int, 0)
	afsafdsasdfdasf(nums, tempAns, &ans)
	return ans
}

func afsafdsasdfdasf(nums []int, tempAns []int, ans *[][]int) {
	if len(nums) == len(tempAns) {
		*ans = append(*ans, append([]int{}, tempAns...))
		return
	}

	// isExist := func(nums []int, val int) bool {
	// 	flg := false
	// 	for i := 0; i < len(nums); i++ {
	// 		if nums[i] == val {
	// 			flg = true
	// 			break
	// 		}
	// 	}
	// 	return flg
	// }

	for i := 0; i < len(nums); i++ {
		// if isExist(tempAns, nums[i]) {
		// 	continue
		// }
		fmt.Println("i ", i)
		tempAns = append(tempAns, nums[i])
		afsafdsasdfdasf(nums, tempAns, ans)
		tempAns = tempAns[:len(tempAns)-1]
	}
}
