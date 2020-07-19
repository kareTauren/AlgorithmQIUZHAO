package most_water

import (
	"fmt"
	"testing"
)

func TestMostWater(t *testing.T) {
	// {1, 8, 6, 2, 5, 4, 8, 3, 7}
	s := []int{0, 1, 0, 3, 12}

	max := func(n, m int) int {
		if n > m {
			return n
		}
		return m
	}

	min := func(n, m int) int {
		if n > m {
			return m
		}
		return n
	}

	// 0, 1, 0, 3, 12
	size := len(s)
	res := 0
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			area := (j - i) * min(s[i], s[j])
			fmt.Println(area)
			res = max(area, res)
		}
	}
	fmt.Println(res)
}

func TestMaxArea(t *testing.T) {
	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(s))
	s = []int{1, 1}
	fmt.Println(maxArea(s))
}

func maxArea(height []int) int {
	max := func(n, m int) int {
		if n > m {
			return n
		}
		return m
	}

	min := func(n, m int) int {
		if n > m {
			return m
		}
		return n
	}
	ans := 0
	left := 0
	right := len(height) - 1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

// func max(n, m int) int {
// 	if n > m {
// 		return n
// 	}
// 	return m
// }
//
// func min(n, m int) int {
// 	if n > m {
// 		return m
// 	}
// 	return n
// }
