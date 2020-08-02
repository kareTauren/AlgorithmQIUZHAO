package home_work

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	fmt.Println(canJump([]int{2, 4, 1, 1, 4}))
}

func canJump(nums []int) bool {
	max := func(n, m int) int {
		if n > m {
			return n
		}
		return m
	}

	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		k = max(k, i+nums[i])
	}
	return true
}
