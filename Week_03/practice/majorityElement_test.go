package practice

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	// fmt.Println("ans ", majorityElement([]int{1, 2, 3, 3}))
	// fmt.Println("ans2 ", majorityElement2([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println("ans3 ", majorityElement3([]int{1, 2, 3, 3, 3}))
	fmt.Println("ans4 ", majorityElement4([]int{1, 2, 3, 3}))

}

// 使用 map 缓存，然后判断
func majorityElement(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	halfLen := len(nums)/2 + 1
	cache := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[nums[i]]++
		if cache[nums[i]] >= halfLen {
			return nums[i]
		}
	}
	return 0
}

func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 随机化 这个方法如果一直找不到数量满足 len(nums)/2 的值，就会一直循环，需要搞一个计数机制，让他停止
func majorityElement3(nums []int) int {

	randomFunc := func(min, max int) int {
		fmt.Println("entry random")
		return rand.Intn(max-min) + min
	}

	halfLen := len(nums) / 2

	findResult := func(nums []int, num int) int {
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] == num {
				count++
			}
		}
		return count
	}

	for {
		randIdx := randomFunc(0, len(nums))
		num := nums[randIdx]
		if findResult(nums, num) > halfLen {
			return num
		}
	}
}

// 投票法
func majorityElement4(nums []int) int {
	rs := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if rs == nums[i] {
			count++
		} else {
			count--
			if count == 0 {
				rs = nums[i]
				count = 1
			}
		}
	}
	return rs
}
