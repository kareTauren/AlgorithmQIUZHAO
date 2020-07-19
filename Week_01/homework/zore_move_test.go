package homework

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	s := []int{1, 1, 0, 3, 12}
	moveZeroes(s)
}

// 时间：O(n)
// 空间：O(1)
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	fmt.Print(nums)
}
