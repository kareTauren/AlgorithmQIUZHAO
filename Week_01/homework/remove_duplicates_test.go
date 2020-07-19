package homework

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicats(t *testing.T) {
	s := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(s))
}

// 快慢双指针
// 时间复杂度：O(N)
// 空间：O(1)
func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
