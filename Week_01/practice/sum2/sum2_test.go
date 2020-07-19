package sum2

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	target := 9
	s := []int{2, 7, 11, 15}
	// fmt.Println(twoSum1(s, target))
	//
	fmt.Println(twoSum2(s, target))
	// fmt.Println(twoSum3(s, target))
}

func twoSum1(nums []int, target int) []int {
	var rs []int
	size := len(nums)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if target == nums[i]+nums[j] {
				rs = append(rs, i, j)
				return rs
			}
		}
	}
	return rs
}

func twoSum2(nums []int, target int) []int {
	rsMap := map[int]int{}
	var rs []int
	for i := 0; i < len(nums); i++ {
		value := target - nums[i]
		if v, ok := rsMap[value]; ok {
			rs = append(rs, v, i)
			break
		}
		rsMap[nums[i]] = i
	}
	return rs
}
