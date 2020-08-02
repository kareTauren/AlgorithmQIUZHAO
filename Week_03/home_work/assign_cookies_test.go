package home_work

import (
	"fmt"
	"sort"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}

// 满足尽可能多的孩子的胃口，那么就要从最小的胃口的孩子派发，这样可以得到尽可能多的人
// 所以需要对孩子胃口数组和饼干数组同时做排序
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}
