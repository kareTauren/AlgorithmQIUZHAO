package sum3

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	s := []int{0, 0, 0}
	// fmt.Println(threeSum_force(s))
	// fmt.Println(hashSolution(s))
	fmt.Println(threeSum(s))
}

// 暴力解法 时间复杂度为 O(n^3)
// 基本思路：三重循环嵌套，然后判断 a + b = -c
// 这个方法都超时了卧槽 垃圾方法
// 用数组去重，效率太低，每次需要遍历 O(n) 复杂度
func threeSum_force(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{}
	}

	// sort.Ints(nums)
	// fmt.Println(nums)

	// 去重函数，看看有没有其他更好的方法😀
	isExist := func(parent [][]int, child []int) bool {
		flag := false
		for _, ary := range parent {
			if reflect.DeepEqual(ary, child) {
				flag = true
				break
			}
		}
		return flag
	}

	var rs [][]int

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					if ok := isExist(rs, []int{nums[i], nums[j], nums[k]}); !ok {
						rs = append(rs, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return rs
}

// 依然要解决去重
func hashSolution(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	var rs [][]int
	for i := 0; i < len(nums)-2; i++ {
		target := -nums[i]
		rsMap := make(map[int]int, len(nums)-i)
		for j := i + 1; j < len(nums); j++ {
			v := target - nums[j]
			if v, ok := rsMap[v]; ok {
				rs = append(rs, []int{nums[i], v, nums[j]})
			} else {
				rsMap[nums[j]] = nums[j]
			}
		}
	}
	return rs
}

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0, 0)
	if len(nums) < 3 {
		return ret
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return ret
}

// 参考
// https://leetcode-cn.com/problems/3sum/solution/san-shu-zhi-he-javajian-ji-ti-jie-by-wang-zi-hao-z/
// https://leetcode-cn.com/problems/3sum/solution/3sumpai-xu-shuang-zhi-zhen-yi-dong-by-jyd/
