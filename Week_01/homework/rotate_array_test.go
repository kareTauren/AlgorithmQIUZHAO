package homework

import (
	"fmt"
	"testing"
)

func TestRotate(T *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	k := 1
	// rotate(s, k)
	rotate2(s, k)
}

// 硬移动 空间复杂度为 k 次 * O(n) , n 是元素个数
// 时间：O(K * N)
// 空间：O(1)
func rotate(nums []int, k int) {
	if len(nums) <= k {
		return
	}

	for i := 0; i < k; i++ {
		prev := nums[len(nums)-1] //  每次取数组最后一个元素往前挪动
		for j := 0; j < len(nums); j++ {
			nums[j], prev = prev, nums[j] // 交换数据
		}
	}
	fmt.Println(nums)
}

// 空间：O(1)
// 时间：O(n)
func rotate2(nums []int, k int) {
	if len(nums) <= k {
		return
	}

	reverse := func(ary []int) {
		for i := 0; i < len(ary)/2; i++ {
			ary[i], ary[len(ary)-1-i] = ary[len(ary)-1-i], ary[i]
		}
	}

	// such as input ary is [1, 2, 3, 4, 5]

	// [5, 4, 3, 2, 1]
	reverse(nums)
	// 找出需要反转前多少个元素 k % 5, k == 3 就是需要反转前三个得 [3， 4， 5, 2, 1]
	reverse(nums[0 : k%len(nums)])
	// 再反转 k 个数字以后的数据，变成升序 [3， 4， 5, 1, 2]
	reverse(nums[k%len(nums):])
	fmt.Print(nums)
}
