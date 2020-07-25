package home_work

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	// nums := []int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 6, 7, 7, 7, 7, 7}
	// k := 3
	nums := []int{1, 2}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}

	ans := make([]int, 0)
	if len(nums) == 0 {
		return ans
	}

	// 记录出现的次数
	cache := map[int]int{}
	for i := 0; i < len(nums); i++ {
		cache[nums[i]]++
	}

	// 排序取到前 k 个高频元素
	tempAry := make([][]int, len(nums)+1)
	for k, v := range cache {
		if len(tempAry[v]) == 0 {
			tempAry[v] = []int{}
		}
		tempAry[v] = append(tempAry[v], k)
	}

	// 把数字出现的次数当做数组的下标来表示，那么在从数组尾部取 k 个值就行
	for i := len(tempAry) - 1; i >= 0 && len(ans) < k; i-- {
		if len(tempAry[i]) == 0 {
			continue
		}
		ans = append(ans, tempAry[i]...)
	}

	return ans
}
