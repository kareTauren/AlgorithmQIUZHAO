package zore_move

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	s := []int{0, 1, 0, 3, 12}
	moveZeroes(s)
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	fmt.Println(nums)
}
