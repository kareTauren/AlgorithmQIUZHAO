package home_work

import (
	"fmt"
	"testing"
)

func TestStandardAryBinarySearch(t *testing.T) {
	fmt.Println(standardAryBinarySearch([]int{5, 6, 7, 1, 2, 3, 4}))
}

func standardAryBinarySearch(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
