package practice

import (
	"testing"
)

func TestSubsets(t *testing.T) {
	// fmt.Println(subsets([]int{1, 2, 3}))
	// fmt.Println()
	// fmt.Println(subsets2([]int{1, 2, 3}))
	// subsets2([]int{1, 2, 3})
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{})

	for i := 0; i < len(nums); i++ {
		newSub := make([][]int, 0)
		for j := 0; j < len(ans); j++ {
			temp := make([]int, 0)
			temp = append(temp, ans[j]...)
			temp = append(temp, nums[i])
			newSub = append(newSub, temp)
		}
		ans = append(ans, newSub...)
	}
	return ans
}

func subsets2(nums []int) [][]int {
	ans := make([][]int, 0)
	tempAns := make([]int, 0)
	backTrack(0, nums, &ans, tempAns)
	return ans
}

func backTrack(n int, num []int, ans *[][]int, tempAns []int) {
	*ans = append(*ans, append([]int{}, tempAns...))
	for i := n; i < len(num); i++ {
		tempAns = append(tempAns, num[i])
		backTrack(i+1, num, ans, tempAns)
		tempAns = tempAns[:len(tempAns)-1]
	}
}
