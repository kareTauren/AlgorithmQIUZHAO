package home_work

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	tempAns := make([]int, 0)
	backTrack(nums, tempAns, &ans)
	return ans
}

func backTrack(nums []int, tempAns []int, ans *[][]int) {
	if len(nums) == len(tempAns) {
		*ans = append(*ans, append([]int{}, tempAns...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if isExist(tempAns, nums[i]) {
			continue
		}
		tempAns = append(tempAns, nums[i])
		backTrack(nums, tempAns, ans)
		tempAns = tempAns[:len(tempAns)-1]
	}
}

func isExist(nums []int, val int) bool {
	flg := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			flg = true
			break
		}
	}
	return flg
}
