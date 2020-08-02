package home_work

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}

func findMin(nums []int) int {
	ans := 0
	l := 0
	r := len(nums) - 1
	target := 0

	for l <= r {
		mid := (l + r) / 2

		if nums[l] > nums[r] {
			l = mid
			target = nums[r]
		} else {
			if nums[mid] < target {
				r = mid - 1
			} else {
				target = nums[mid]
				l = mid + 1
			}
		}
	}
	return ans
}
