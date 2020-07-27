package home_work

import (
	"fmt"
	"sort"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
}

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	tempAns := make([]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backTrack2(nums, tempAns, &ans, used)
	return ans
}

func backTrack2(nums []int, tempAns []int, ans *[][]int, used []bool) {
	if len(nums) == len(tempAns) {
		*ans = append(*ans, append([]int{}, tempAns...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		tempAns = append(tempAns, nums[i])
		backTrack2(nums, tempAns, ans, used)
		tempAns = tempAns[:len(tempAns)-1]
		used[i] = false
	}
}
