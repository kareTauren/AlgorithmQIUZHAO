package homework

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{0}
	nums2 := []int{1}
	m := 0
	n := 1
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy_num1 := func(num []int, m int) []int {
		var res []int
		for i := 0; i < m; i++ {
			res = append(res, num[i])
		}
		return res
	}

	subNums1 := copy_num1(nums1, m)

	p1 := 0
	p2 := 0
	idx := 0
	for p1 < m && p2 < n {
		if subNums1[p1] <= nums2[p2] {
			nums1[idx] = subNums1[p1]
			p1++
		} else {
			nums1[idx] = nums2[p2]
			p2++
		}
		idx++
	}
	if p1 < m {
		for p1 < m {
			nums1[idx] = subNums1[p1]
			p1++
			idx++
		}
	}
	if p2 < n {
		for p2 < n {
			nums1[idx] = nums2[p2]
			p2++
			idx++
		}
	}

	fmt.Println(nums1)
}
