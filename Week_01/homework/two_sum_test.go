package homework

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	s := []int{2, 3, 4, 6}
	target := 10
	fmt.Println(twoSum2(s, target))
}

// 空间：O(n)
// 时间：O(n)
// cache
func twoSum(nums []int, target int) []int {
	var res []int
	cache := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		s := target - nums[i]
		if v, ok := cache[s]; ok {
			res = append(res, v, i)
			break
		}
		cache[nums[i]] = i
	}
	return res
}

// 空间：O(1)
// 时间：O(n*n)
func twoSum2(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				res = append(res, i, j)
				break
			}
		}
	}
	return res
}
