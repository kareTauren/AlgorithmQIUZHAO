package home_work

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	fmt.Println(majorityElement([]int{1, 2, 2, 2, 2}))
}

// 1 大于 ⌊ n/2 ⌋
func majorityElement(nums []int) int {
	half := len(nums) / 2
	numCache := map[int]int{}
	for i := 0; i < len(nums); i++ {
		numCache[nums[i]]++
		if numCache[nums[i]] > half {
			return nums[i]
		}
	}
	return 0
}
